# selfX
selfX is a simple and lightweight web content management system

------


## Get started
+ [Download program file](https://github.com/icoller/selfx/releases)
+ Run

      ./selfx

+ Start successfully
> Use sqlite by default<br>
> Default management path /admin

------

### Configuration file(conf.toml)

| key  | Description       | Default   |
|------|-------------------|-----------|
| addr | listening address | random    |
| db   | database type     | sqlite    |
| dsn  | data source name  | ./selfx.db |

+ Data source name Examples

| Type       | dsn Example                                                                        |
|------------|------------------------------------------------------------------------------------|
| sqlite     | ./data.db                                                                          |
| mysql      | user:password@tcp(127.0.0.1:3306)/selfx?charset=utf8mb4&parseTime=True              |
| postgresql | host=127.0.0.1 port=5432 user=postgres password=123456 dbname=selfx sslmode=disable |



### command
| key         | Description       | Example                                  |
|-------------|----------|----------------------------------------|
| --username  | reset administrator username |                                        |
| --password  | reset administrator password  |                                        |
| --adminpath | reset administration path    | ./selfx --adminpath="admin"             |
| --config    | Define profile path | ./selfx --config="/home/othername.toml" |

> ###### show more information by ./selfx --help 