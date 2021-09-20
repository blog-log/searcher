package adapter

import "bloglog.com/search/v1/internal/model"

// convert page object to algolia object
func PageToAlgolia(page *model.Page) *model.AlgoliaPage {
	return &model.AlgoliaPage{
		Id:      page.Id,
		Repo:    page.Repo,
		Branch:  page.Branch,
		Path:    page.Path,
		Title:   page.Title,
		Content: page.Content,
	}
}

// convert algolia object to page object
func AlgoliaToPage(algoliaPage *model.AlgoliaPage) *model.Page {
	return &model.Page{
		Id:      algoliaPage.Id,
		Repo:    algoliaPage.Repo,
		Branch:  algoliaPage.Branch,
		Path:    algoliaPage.Path,
		Title:   algoliaPage.Title,
		Content: algoliaPage.Content,
	}
}
