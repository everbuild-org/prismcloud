package services

import (
	"prismcloud.dev/apiserver/database/model"
)

func (s *ServiceApi) HasNamespace(name string) bool {
	namespace := &model.Namespace{}
	err := s.Database.First(&namespace, "name = ?", name).Error
	return err == nil
}

func (s *ServiceApi) CreateNamespace(name string, ramLimit int64) error {
	namespace := &model.Namespace{
		Name:     name,
		RamLimit: ramLimit,
	}

	if err := s.Database.Create(namespace).Error; err != nil {
		return err
	}

	return nil
}

func (s *ServiceApi) DeleteNamespace(name string) error {
	if err := s.Database.Where("name = ?", name).Delete(&model.Namespace{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *ServiceApi) GetNamespace(name string) (*model.Namespace, error) {
	namespace := &model.Namespace{}
	if err := s.Database.Where("name = ?", name).First(&namespace).Error; err != nil {
		return nil, err
	}

	return namespace, nil
}

func (s *ServiceApi) ListNamespaces() ([]*model.Namespace, error) {
	var namespaces []*model.Namespace
	if err := s.Database.Find(&namespaces).Error; err != nil {
		return nil, err
	}

	return namespaces, nil
}

func (s *ServiceApi) PatchNamespace(name string, ramLimit int64) error {
	namespace, err := s.GetNamespace(name)

	if err != nil {
		return err
	}

	namespace.RamLimit = ramLimit

	if err := s.Database.Save(namespace).Error; err != nil {
		return err
	}

	return nil
}
