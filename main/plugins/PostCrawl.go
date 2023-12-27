package plugins

import (
	"fmt"
	apiServ "selfx/app/api/service"
	"selfx/app/model"
	pluginEntity "selfx/app/plugin/entity"
	"selfx/app/repo/context"
	"selfx/app/service"
	"selfx/init/db"
	"sync"

	"go.uber.org/zap"
)

type PostCrawl struct {
	Limit           int   `json:"limit"`
	Order           int   `json:"order"`             // 0:最新 1:最早 2:随机
	CategoryIds     []int `json:"category_ids"`      // 指定栏目ID
	DeleteOnFailure bool  `json:"delete_on_failure"` // 失败时删除
	ctx             *pluginEntity.Plugin
}

func NewPostCrawl() *PostCrawl {
	return &PostCrawl{
		Limit:           1,
		Order:           1,
		DeleteOnFailure: true,
		CategoryIds:     []int{},
	}
}

func (p *PostCrawl) Info() *pluginEntity.PluginInfo {
	return &pluginEntity.PluginInfo{
		ID:         "PostCrawl",
		About:      "publish articles form Crawl",
		RunEnable:  true,
		CronEnable: true,
		PluginInfoPersistent: pluginEntity.PluginInfoPersistent{
			CronStart: false,
			CronExp:   "@every 1h",
		},
	}
}

func (p *PostCrawl) Load(ctx *pluginEntity.Plugin) error {
	return nil
}

func (p *PostCrawl) Run(ctx *pluginEntity.Plugin) (err error) {
	p.ctx = ctx

	if p.Limit <= 0 {
		p.ctx.Log.Warn("limit <= 0")
		return
	}
	ids, err := p.ids()
	if err != nil {
		return
	}
	if len(ids) == 0 {
		if len(p.CategoryIds) > 0 {
			p.ctx.Log.Warn("NotFound Articles")
		}
		return
	}
	var success int
	var wg = &sync.WaitGroup{}
	for _, id := range ids {
		wg.Add(1)
		go p.post(id, wg, &success)
	}
	wg.Wait()
	p.ctx.Log.Info(fmt.Sprintf("End. success total: %d failures: %d", success, len(ids)-success))
	return nil
}

func (p *PostCrawl) post(id int, wg *sync.WaitGroup, success *int) {
	defer wg.Done()
	item, err := apiServ.CrawlPost(id)
	if err != nil {
		p.ctx.Log.Error("post error", zap.Error(err), zap.Int("id", id))
		if p.DeleteOnFailure {
			if e := service.Crawl.Delete(id); e != nil {
				p.ctx.Log.Error("delete error", zap.Error(err), zap.Int("id", id))
			}
		}
		return
	}
	*success++
	p.ctx.Log.Info(fmt.Sprintf("Post success, ID: %d Title: %s", item.ID, item.Title))
	return
}

func (p *PostCrawl) ids() (res []int, err error) {
	var items []model.Crawl
	var ctx = context.Context{Select: "id", Limit: p.Limit, Order: p.order()}
	if len(p.CategoryIds) > 0 {
		items, err = service.Crawl.ListByCategoryIds(&ctx, p.CategoryIds)
	} else {
		items, err = service.Crawl.List(&ctx)
	}
	if err != nil {
		p.ctx.Log.Error("query list error", zap.Error(err))
		return
	}
	for _, item := range items {
		res = append(res, item.ID)
	}
	return
}

func (p *PostCrawl) order() any {
	switch p.Order {
	case 2:
		return db.RandomOrder()
	case 1:
		return "id asc"
	default:
		return "id desc"
	}
}
