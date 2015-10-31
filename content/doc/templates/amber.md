---
aliases:
- /doc/templates/amber
- /doc/layout/templates/amber
- /doc/layout/amber/
date: 2015-07-20
linktitle: Amber templates
menu:
  doc:
    parent: layout
next: /doc/templates/functions
prev: /doc/templates/go-templates
title: Amber Templates
weight: 18
---

**Note:** The Amber support is broken in Hugo 0.14, but is fixed in the upcoming release.

Amber templates are another template type which Hugo supports, in addition to [Go templates](/doc/templates/go-templates) and [Ace templates](/doc/templates/ace-templates) templates.

For template documentation, follow the links from the [Amber project](https://github.com/eknkc/amber)

* Amber templates must be named with the amber-suffix, e.g. `list.amber`
* Partials in Amber or HTML can be included with the Amber template syntax:
	* `import ../partials/test.html `
	* `import ../partials/test_a.amber `


