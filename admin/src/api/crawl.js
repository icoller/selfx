import { usePost } from "./hook";

export const crawlRuleAdd = (data) => usePost("/crawl/rule/add", data);

