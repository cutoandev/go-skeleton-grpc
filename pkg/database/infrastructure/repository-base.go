package infrastructure

import (
  "gorm.io/gorm"
)

type RepositoryBase[K any] struct {
  DB        *gorm.DB
  modelType K
}

// DBContext return the database context
//  Output:
//    + gorm.DB: database context
func (repo RepositoryBase[K]) DBContext() *gorm.DB {
  return repo.DB
}

// Model return database model
//  Output:
//    + gorm.DB: database context with model mapped
func (repo RepositoryBase[K]) Model() *gorm.DB {
  return repo.DB.Model(repo.modelType)
}

// GetById this func will query data model by ID
//  In parameters:
//    + id: key of model to query data
//  Out data:
//    + K: the model found
//    + error: error info, nil if there is no error
func (repo RepositoryBase[K]) GetById(id interface{}) (K, error) {
  database := repo.DB
  var result K
  if err := database.Model(repo.modelType).Find(&result, id).Error; err != nil {
    return result, err
  }
  return result, nil
}

// GetAll this func will return all data from database
//  In parameters: no input parameter(s)
//  Out data:
//    + []K: the list of data found
//    + error: error info, nil if there is no error
func (repo RepositoryBase[K]) GetAll() ([]K, error) {
  database := repo.DB
  var result []K
  if err := database.Model(repo.modelType).Find(&result).Error; err != nil {
    return result, err
  }
  return result, nil
}

// Where this function will query on table model with the condition for Where clause
//  In parameters:
//    + condition: where clause, unlimited condition
//  Out data:
//    + []K: the list of data found
//    + error: error info, nil if there is no error
func (repo RepositoryBase[K]) Where(condition ...any) ([]K, error) {
  database := repo.DB
  var result []K
  if err := database.Model(repo.modelType).Where(condition).Find(&result).Error; err != nil {
    return result, err
  }
  return result, nil
}

// Create is the function for create new record to database
//  In parameters:
//    + entity: model data to update
//    + tx: transaction for control rollback when having error. Input nil if no need
//  Out data:
//    + error: error information, nil if there is no error
func (repo RepositoryBase[K]) Create(entity K, tx *gorm.DB) error {
  database := repo.DB
  if tx != nil {
    database = tx
  }

  defer func() {
    if r := recover(); r != nil {
      database.Rollback()
    }
  }()

  if err := database.Error; err != nil {
    return err
  }

  if err := database.Model(repo.modelType).Create(entity).Error; err != nil {
    database.Rollback()
    return err
  }

  return nil
}

// Update is the function for update multiple columns to database
//  In parameters:
//    + entity: model data to update
//    + tx: transaction for control rollback when having error. Input nil if no need
//    + condition: condition to update. Follow rules of GORM
//  Out data:
//    + error: error information, nil if there is no error
func (repo RepositoryBase[K]) Update(entity K, tx *gorm.DB, condition ...any) error {
  database := repo.DB
  if tx != nil {
    database = tx
  }

  defer func() {
    if r := recover(); r != nil {
      database.Rollback()
    }
  }()

  if err := database.Error; err != nil {
    return err
  }

  if err := database.Model(repo.modelType).Where(condition).Updates(entity).Error; err != nil {
    database.Rollback()
    return err
  }

  return nil
}

// Delete is the function for delete data
//  In parameters:
//    + entity: model data to delete
//    + tx: transaction for control rollback when having error. Input nil if no need
//    + condition: condition to delete. Follow rules of GORM
//  Out data:
//    + error: error information, nil if there is no error
func (repo RepositoryBase[K]) Delete(entity K, tx *gorm.DB, condition ...any) error {
  database := repo.DB
  if tx != nil {
    database = tx
  }

  defer func() {
    if r := recover(); r != nil {
      database.Rollback()
    }
  }()

  if err := database.Error; err != nil {
    return err
  }

  if err := database.Model(repo.modelType).Where(condition).Delete(entity).Error; err != nil {
    database.Rollback()
    return err
  }

  return nil
}
