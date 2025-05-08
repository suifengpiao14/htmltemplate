增强html
为 HTML 模板中的所有非 <script> 节点设置：
一个唯一的 data-node-key
一个动态属性占位符：{{<nodeKey>Attrs}}=""(html_render.go 会写入 {{<nodeKey>Attrs}})