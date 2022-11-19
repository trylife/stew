# Elasticsearch

# 数据类型

## numeric 
```
long
integrer
short
byte
double
float
half_float
scaled_float
unsigned_long
```

## keywords

## dates

## alias

## binary

## range

```go
integer_range
float_range
long_range
double_range
date_range
```
## text

## object

object

nested

flattened

## 结构化数据类型

geo-point
geo-shape
point
shape

## 特殊类型

```go
ip
completeion
tocken_count
nurmur3
annotated-text
percolator
join
rank_features
dese_vector
sparse_vector
search-as_you-type
histogram
constant_keyword
)

```

## array 

## 新增
date_nanos 纳秒
features

# Mapping

- mapping 映射
    - dynamic mapping 动态映射
        - ?
        - runtime 运行时映射
    - explicit mapping 显示映射

## 映射规则

| JSON数据类型     | "dynamic":"true"  | "dynamic": "runtime" | 
|--------------|-------------------|----------------------|
| true / false | bollean           | boolean              | 
| double       | float             | double               |
| integer      | long              | long                 |
| object       | object            | object               |
| array        | 取决于第一个元素类型        | 取决于第一个元素类型           |
| 日期类型         | date              | date                 |
| 数字类型         | float / long      | double / long        |
| 日期和数字之外的类型   | 带 .keyword 子字段的文本 | keyword      |

## date 日期

- 日期储存
- 禁用日期映射 `"date_detection": false`
- 自定义日期检测格式 `"dynamic_date_formats": ["yyy/MM/dd"]`

## 数字映射

- 默认关闭数字映射检测
- 开启使用数字映射检测 `"numeric_detection": true`

## 动态映射模板

## text
- 用于全文检索
- 搜索前分词器将文字内容转为分词(term)
- text类型不适合排序
- 如果需要排序，推荐使用keyword

## keyword

# 概念
- 倒排索引 `用于文本检索`
- 正排索引 `用于做排序和聚合`
- 

# 常用的数据类型
- text `当一个字段被全文检索; 不能用做聚合`
- keyword `不会被分词`
- alias
- object `保存对象`
- nested `存储对象数组`
- geo-point `地理位置-点`
- geo-shape `地理位置图形`
- IP `IP`
- completion `搜索建议`

# Mapping
- mapping智能创建不能修改？
- setting 设置一些选项

## 映射参数
- index 是否对当前字段创建倒排索引
- analyzer 指定分词器
- doc_text `text字段设置正排索引，用于排序和聚合时用`
- fielddata `内存中动态的正排索引，一般禁止使用`

- TODO reindex?

# aggregations 聚合
- _source
- exact_value

# query DSL `domain spec language`

- match

```http request
GET product/_search

{
  "query": {
    "match": {
      "name": "xiaomi nfc phone"
    }
  }
}


```


```http request
// 查询所有
GET product/_search

{
  "query": {
    "match_all": {}
  }
}

```

```http request
// ?a=&b?
GET product/_search

{
  "query": {
    "multi_match": {
    "query": "phone huangmenji",
    "fields": ["name", "desc"],
    }
  }
}

```

- term
- range
- nested
- wildcard