#数据解析器
目前支持xml、json 2种格式，2种格式都支持模板引擎
如果模版中有数组需要循环生成，只能使用xml格式
其它情况建议可以使用json格式
**注意实现**
xml格式需要生成纯数字时，不能携带属性，例如：
```xml
<members>1</members>
<members>2</members>
<members>3</members>
<members>4</members>
```
将生成数据:
```json
{"members":[1,2,3,4]}
```
而
```xml
<members type="array">1</members>
<members type="array">2</members>
<members type="array">3</members>
<members type="array">4</members>
```
将生成数据:
```json
{"members":[
    {"-type":"array","#text":1},
    {"-type":"array","#text":2},
    {"-type":"array","#text":3},
    {"-type":"array","#text":4},
]}
```