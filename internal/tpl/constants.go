package tpl

const TplTag = `---
title: Posts para el tag {{ .tag }}
---

{{ range .posts }}
<section>
	## [{{ .title }}](/posts/{{ .slug }})

	*{{.desc}}*
	<time>{{ .date }}</time>
</section>
{{ end }}
`

const AllTags = `---
title: Listado de tags
---


{{ range .tags }}
-   [{{.Tag}}](/tag/{{.Slug}})
{{- end }}
`

const AllPosts = `---
title: Listado de posts
---

{{ range .posts }}
<section>
<time>{{ .date }}</time>
[**{{ .title }}**](/post/{{ .slug }})
*{{.desc}}*
</section>
{{ end }}`
