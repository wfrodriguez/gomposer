package tpl

const TplTag = `# Posts para el tag **{{.tag}}**

{{ range .posts }}
<section>
	## [{{ .title }}](/posts/{{ .slug }})

	*{{.desc}}*
	<time>{{ .date }}</time>
</section>
{{ end }}
`
