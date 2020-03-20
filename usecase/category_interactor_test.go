package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
	"testing"
)

type MockCategoryRepository struct {
}

func (repo *MockCategoryRepository) All() (categories model.Categories, err error) {
	categories = model.Categories{
		model.Category{Identifier: "test", JpName: "test", Color: "#000000"},
		model.Category{Identifier: "test", JpName: "test", Color: "#000000"},
		model.Category{Identifier: "test", JpName: "test", Color: "#000000"},
		model.Category{Identifier: "test", JpName: "test", Color: "#000000"},
		model.Category{Identifier: "test", JpName: "test", Color: "#000000"},
		model.Category{Identifier: "test", JpName: "test", Color: "#000000"},
		model.Category{Identifier: "test", JpName: "test", Color: "#000000"},
		model.Category{Identifier: "test", JpName: "test", Color: "#000000"},
		model.Category{Identifier: "test", JpName: "test", Color: "#000000"},
		model.Category{Identifier: "test", JpName: "test", Color: "#000000"},
	}
	return categories, nil
}

func (repo *MockCategoryRepository) FindByIdentifier(string) (category model.Category, err error) {
	category = model.Category{
		Identifier: "test",
		JpName:     "test",
		Color:      "#000000",
	}
	return category, nil
}

func (repo *MockCategoryRepository) Create(repository.CategoryCreateParams) error {
	return nil
}

func (repo *MockCategoryRepository) Update(repository.CategoryCreateParams) error {
	return nil
}

func (repo *MockCategoryRepository) Delete(string) error {
	return nil
}

func TestShouldCategoriesIndex(t *testing.T) {
	interactor := initTestCategoryInteractor()

	t.Run("Successful Request", func(t *testing.T) {
		categories, err := interactor.Index()
		if err != nil {
			t.Fatalf("Cannot get categories: %s", err)
		}
		if len(categories) != 10 {
			t.Fatalf("Fail expected: 10, got: %v", len(categories))
		}
	})
}

func TestShouldFindCategoryByIdentifier(t *testing.T) {
	interactor := initTestCategoryInteractor()
	category, err := interactor.FindByIdentifier("test")
	if err != nil {
		t.Fatalf("Cannot get category: %s", err)
	}
	if category.Identifier != "test" {
		t.Fatalf("Fail expected identifier: test, got: %v", category)
	}
}

func TestShouldCreateCategory(t *testing.T) {
	interactor := initTestCategoryInteractor()
	params := CategoryCreateParams{
		Identifier: "test",
		JpName:     "test",
		Color:      "#000000",
	}

	err := interactor.Create(params)
	if err != nil {
		t.Fatalf("Could not create category: %s", err.Error())
	}

	err = interactor.Update(params)
	if err != nil {
		t.Fatalf("fail to update category: %s", err.Error())
	}
}

func TestShouldDeleteCategory(t *testing.T) {
	interactor := initTestCategoryInteractor()
	err := interactor.Delete("test")
	if err != nil {
		t.Fatal("fail to delete")
	}
}

func initTestCategoryInteractor() CategoryInteractor {
	return CategoryInteractor{
		CategoryRepository: new(MockCategoryRepository),
	}
}
