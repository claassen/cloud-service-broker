// Copyright 2019 the Service Broker Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go generate; DO NOT EDIT.

package db_service

import (
	"context"

	"github.com/GoogleCloudPlatform/gcp-service-broker/db_service/models"
	"github.com/jinzhu/gorm"
)

// CreateServiceInstanceDetails creates a new record in the database and assigns it a primary key.
func CreateServiceInstanceDetails(ctx context.Context, object *models.ServiceInstanceDetails) error { return defaultDatastore().CreateServiceInstanceDetails(ctx, object) }
func (ds *SqlDatastore) CreateServiceInstanceDetails(ctx context.Context, object *models.ServiceInstanceDetails) error {
	return ds.db.Create(object).Error
}

// SaveServiceInstanceDetails updates an existing record in the database.
func SaveServiceInstanceDetails(ctx context.Context, object *models.ServiceInstanceDetails) error { return defaultDatastore().SaveServiceInstanceDetails(ctx, object) }
func (ds *SqlDatastore) SaveServiceInstanceDetails(ctx context.Context, object *models.ServiceInstanceDetails) error {
	return ds.db.Save(object).Error
}
// DeleteServiceInstanceDetailsById soft-deletes the record by its key (id).
func DeleteServiceInstanceDetailsById(ctx context.Context, id string) error { return defaultDatastore().DeleteServiceInstanceDetailsById(ctx, id) }
func (ds *SqlDatastore) DeleteServiceInstanceDetailsById(ctx context.Context, id string) error {
	return ds.db.Where("id = ?", id).Delete(&models.ServiceInstanceDetails{}).Error
}



// DeleteServiceInstanceDetails soft-deletes the record.
func DeleteServiceInstanceDetails(ctx context.Context, record *models.ServiceInstanceDetails) error { return defaultDatastore().DeleteServiceInstanceDetails(ctx, record) }
func (ds *SqlDatastore) DeleteServiceInstanceDetails(ctx context.Context, record *models.ServiceInstanceDetails) error {
	return ds.db.Delete(record).Error
}
// GetServiceInstanceDetailsById gets an instance of ServiceInstanceDetails by its key (id).
func GetServiceInstanceDetailsById(ctx context.Context, id string) (*models.ServiceInstanceDetails, error) { return defaultDatastore().GetServiceInstanceDetailsById(ctx, id) }
func (ds *SqlDatastore) GetServiceInstanceDetailsById(ctx context.Context, id string) (*models.ServiceInstanceDetails, error) {
	record := models.ServiceInstanceDetails{}
	if err := ds.db.Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

// ExistsServiceInstanceDetailsById checks to see if an instance of ServiceInstanceDetails exists by its key (id).
func ExistsServiceInstanceDetailsById(ctx context.Context, id string) (bool, error) { return defaultDatastore().ExistsServiceInstanceDetailsById(ctx, id) }
func (ds *SqlDatastore) ExistsServiceInstanceDetailsById(ctx context.Context, id string) (bool, error) {
	return recordToExists(ds.GetServiceInstanceDetailsById(ctx, id))
}



// CreateServiceBindingCredentials creates a new record in the database and assigns it a primary key.
func CreateServiceBindingCredentials(ctx context.Context, object *models.ServiceBindingCredentials) error { return defaultDatastore().CreateServiceBindingCredentials(ctx, object) }
func (ds *SqlDatastore) CreateServiceBindingCredentials(ctx context.Context, object *models.ServiceBindingCredentials) error {
	return ds.db.Create(object).Error
}

// SaveServiceBindingCredentials updates an existing record in the database.
func SaveServiceBindingCredentials(ctx context.Context, object *models.ServiceBindingCredentials) error { return defaultDatastore().SaveServiceBindingCredentials(ctx, object) }
func (ds *SqlDatastore) SaveServiceBindingCredentials(ctx context.Context, object *models.ServiceBindingCredentials) error {
	return ds.db.Save(object).Error
}
// DeleteServiceBindingCredentialsByServiceInstanceIdAndBindingId soft-deletes the record by its key (serviceInstanceId, bindingId).
func DeleteServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx context.Context, serviceInstanceId string, bindingId string) error { return defaultDatastore().DeleteServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx, serviceInstanceId, bindingId) }
func (ds *SqlDatastore) DeleteServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx context.Context, serviceInstanceId string, bindingId string) error {
	return ds.db.Where("service_instance_id = ? AND binding_id = ?", serviceInstanceId, bindingId).Delete(&models.ServiceBindingCredentials{}).Error
}

// DeleteServiceBindingCredentialsByBindingId soft-deletes the record by its key (bindingId).
func DeleteServiceBindingCredentialsByBindingId(ctx context.Context, bindingId string) error { return defaultDatastore().DeleteServiceBindingCredentialsByBindingId(ctx, bindingId) }
func (ds *SqlDatastore) DeleteServiceBindingCredentialsByBindingId(ctx context.Context, bindingId string) error {
	return ds.db.Where("binding_id = ?", bindingId).Delete(&models.ServiceBindingCredentials{}).Error
}

// DeleteServiceBindingCredentialsById soft-deletes the record by its key (id).
func DeleteServiceBindingCredentialsById(ctx context.Context, id uint) error { return defaultDatastore().DeleteServiceBindingCredentialsById(ctx, id) }
func (ds *SqlDatastore) DeleteServiceBindingCredentialsById(ctx context.Context, id uint) error {
	return ds.db.Where("id = ?", id).Delete(&models.ServiceBindingCredentials{}).Error
}



// DeleteServiceBindingCredentials soft-deletes the record.
func DeleteServiceBindingCredentials(ctx context.Context, record *models.ServiceBindingCredentials) error { return defaultDatastore().DeleteServiceBindingCredentials(ctx, record) }
func (ds *SqlDatastore) DeleteServiceBindingCredentials(ctx context.Context, record *models.ServiceBindingCredentials) error {
	return ds.db.Delete(record).Error
}
// GetServiceBindingCredentialsByServiceInstanceIdAndBindingId gets an instance of ServiceBindingCredentials by its key (serviceInstanceId, bindingId).
func GetServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx context.Context, serviceInstanceId string, bindingId string) (*models.ServiceBindingCredentials, error) { return defaultDatastore().GetServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx, serviceInstanceId, bindingId) }
func (ds *SqlDatastore) GetServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx context.Context, serviceInstanceId string, bindingId string) (*models.ServiceBindingCredentials, error) {
	record := models.ServiceBindingCredentials{}
	if err := ds.db.Where("service_instance_id = ? AND binding_id = ?", serviceInstanceId, bindingId).First(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

// ExistsServiceBindingCredentialsByServiceInstanceIdAndBindingId checks to see if an instance of ServiceBindingCredentials exists by its key (serviceInstanceId, bindingId).
func ExistsServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx context.Context, serviceInstanceId string, bindingId string) (bool, error) { return defaultDatastore().ExistsServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx, serviceInstanceId, bindingId) }
func (ds *SqlDatastore) ExistsServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx context.Context, serviceInstanceId string, bindingId string) (bool, error) {
	return recordToExists(ds.GetServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx, serviceInstanceId, bindingId))
}

// GetServiceBindingCredentialsByBindingId gets an instance of ServiceBindingCredentials by its key (bindingId).
func GetServiceBindingCredentialsByBindingId(ctx context.Context, bindingId string) (*models.ServiceBindingCredentials, error) { return defaultDatastore().GetServiceBindingCredentialsByBindingId(ctx, bindingId) }
func (ds *SqlDatastore) GetServiceBindingCredentialsByBindingId(ctx context.Context, bindingId string) (*models.ServiceBindingCredentials, error) {
	record := models.ServiceBindingCredentials{}
	if err := ds.db.Where("binding_id = ?", bindingId).First(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

// ExistsServiceBindingCredentialsByBindingId checks to see if an instance of ServiceBindingCredentials exists by its key (bindingId).
func ExistsServiceBindingCredentialsByBindingId(ctx context.Context, bindingId string) (bool, error) { return defaultDatastore().ExistsServiceBindingCredentialsByBindingId(ctx, bindingId) }
func (ds *SqlDatastore) ExistsServiceBindingCredentialsByBindingId(ctx context.Context, bindingId string) (bool, error) {
	return recordToExists(ds.GetServiceBindingCredentialsByBindingId(ctx, bindingId))
}

// GetServiceBindingCredentialsById gets an instance of ServiceBindingCredentials by its key (id).
func GetServiceBindingCredentialsById(ctx context.Context, id uint) (*models.ServiceBindingCredentials, error) { return defaultDatastore().GetServiceBindingCredentialsById(ctx, id) }
func (ds *SqlDatastore) GetServiceBindingCredentialsById(ctx context.Context, id uint) (*models.ServiceBindingCredentials, error) {
	record := models.ServiceBindingCredentials{}
	if err := ds.db.Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

// ExistsServiceBindingCredentialsById checks to see if an instance of ServiceBindingCredentials exists by its key (id).
func ExistsServiceBindingCredentialsById(ctx context.Context, id uint) (bool, error) { return defaultDatastore().ExistsServiceBindingCredentialsById(ctx, id) }
func (ds *SqlDatastore) ExistsServiceBindingCredentialsById(ctx context.Context, id uint) (bool, error) {
	return recordToExists(ds.GetServiceBindingCredentialsById(ctx, id))
}



// CreateProvisionRequestDetails creates a new record in the database and assigns it a primary key.
func CreateProvisionRequestDetails(ctx context.Context, object *models.ProvisionRequestDetails) error { return defaultDatastore().CreateProvisionRequestDetails(ctx, object) }
func (ds *SqlDatastore) CreateProvisionRequestDetails(ctx context.Context, object *models.ProvisionRequestDetails) error {
	return ds.db.Create(object).Error
}

// SaveProvisionRequestDetails updates an existing record in the database.
func SaveProvisionRequestDetails(ctx context.Context, object *models.ProvisionRequestDetails) error { return defaultDatastore().SaveProvisionRequestDetails(ctx, object) }
func (ds *SqlDatastore) SaveProvisionRequestDetails(ctx context.Context, object *models.ProvisionRequestDetails) error {
	return ds.db.Save(object).Error
}
// DeleteProvisionRequestDetailsById soft-deletes the record by its key (id).
func DeleteProvisionRequestDetailsById(ctx context.Context, id uint) error { return defaultDatastore().DeleteProvisionRequestDetailsById(ctx, id) }
func (ds *SqlDatastore) DeleteProvisionRequestDetailsById(ctx context.Context, id uint) error {
	return ds.db.Where("id = ?", id).Delete(&models.ProvisionRequestDetails{}).Error
}



// DeleteProvisionRequestDetails soft-deletes the record.
func DeleteProvisionRequestDetails(ctx context.Context, record *models.ProvisionRequestDetails) error { return defaultDatastore().DeleteProvisionRequestDetails(ctx, record) }
func (ds *SqlDatastore) DeleteProvisionRequestDetails(ctx context.Context, record *models.ProvisionRequestDetails) error {
	return ds.db.Delete(record).Error
}
// GetProvisionRequestDetailsById gets an instance of ProvisionRequestDetails by its key (id).
func GetProvisionRequestDetailsById(ctx context.Context, id uint) (*models.ProvisionRequestDetails, error) { return defaultDatastore().GetProvisionRequestDetailsById(ctx, id) }
func (ds *SqlDatastore) GetProvisionRequestDetailsById(ctx context.Context, id uint) (*models.ProvisionRequestDetails, error) {
	record := models.ProvisionRequestDetails{}
	if err := ds.db.Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

// ExistsProvisionRequestDetailsById checks to see if an instance of ProvisionRequestDetails exists by its key (id).
func ExistsProvisionRequestDetailsById(ctx context.Context, id uint) (bool, error) { return defaultDatastore().ExistsProvisionRequestDetailsById(ctx, id) }
func (ds *SqlDatastore) ExistsProvisionRequestDetailsById(ctx context.Context, id uint) (bool, error) {
	return recordToExists(ds.GetProvisionRequestDetailsById(ctx, id))
}



// CreateTerraformDeployment creates a new record in the database and assigns it a primary key.
func CreateTerraformDeployment(ctx context.Context, object *models.TerraformDeployment) error { return defaultDatastore().CreateTerraformDeployment(ctx, object) }
func (ds *SqlDatastore) CreateTerraformDeployment(ctx context.Context, object *models.TerraformDeployment) error {
	return ds.db.Create(object).Error
}

// SaveTerraformDeployment updates an existing record in the database.
func SaveTerraformDeployment(ctx context.Context, object *models.TerraformDeployment) error { return defaultDatastore().SaveTerraformDeployment(ctx, object) }
func (ds *SqlDatastore) SaveTerraformDeployment(ctx context.Context, object *models.TerraformDeployment) error {
	return ds.db.Save(object).Error
}
// DeleteTerraformDeploymentById soft-deletes the record by its key (id).
func DeleteTerraformDeploymentById(ctx context.Context, id string) error { return defaultDatastore().DeleteTerraformDeploymentById(ctx, id) }
func (ds *SqlDatastore) DeleteTerraformDeploymentById(ctx context.Context, id string) error {
	return ds.db.Where("id = ?", id).Delete(&models.TerraformDeployment{}).Error
}



// DeleteTerraformDeployment soft-deletes the record.
func DeleteTerraformDeployment(ctx context.Context, record *models.TerraformDeployment) error { return defaultDatastore().DeleteTerraformDeployment(ctx, record) }
func (ds *SqlDatastore) DeleteTerraformDeployment(ctx context.Context, record *models.TerraformDeployment) error {
	return ds.db.Delete(record).Error
}
// GetTerraformDeploymentById gets an instance of TerraformDeployment by its key (id).
func GetTerraformDeploymentById(ctx context.Context, id string) (*models.TerraformDeployment, error) { return defaultDatastore().GetTerraformDeploymentById(ctx, id) }
func (ds *SqlDatastore) GetTerraformDeploymentById(ctx context.Context, id string) (*models.TerraformDeployment, error) {
	record := models.TerraformDeployment{}
	if err := ds.db.Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

// ExistsTerraformDeploymentById checks to see if an instance of TerraformDeployment exists by its key (id).
func ExistsTerraformDeploymentById(ctx context.Context, id string) (bool, error) { return defaultDatastore().ExistsTerraformDeploymentById(ctx, id) }
func (ds *SqlDatastore) ExistsTerraformDeploymentById(ctx context.Context, id string) (bool, error) {
	return recordToExists(ds.GetTerraformDeploymentById(ctx, id))
}



func recordToExists(_ interface{}, err error) (bool, error) {
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
