# Kube-Vaccine
![visitor badge](https://visitor-badge.laobi.icu/badge?page_id=shivi28.kube-vaccine)

This project demonstrates how to write API conversions for Kubernetes CRDs by taking a use case of Covid19 Vaccine registration.

------

### Why do we need to do API conversions?


### Write API conversions for kubernetes custom resource "Registration"  

----
### Prerequisite
1. Install Go
2. kind cluster
3. kubebuilder
----
### Steps

#### Step1: Create kubebuilder project
```bash
>> mkdir kube-vaccine
>> cd kube-vaccine
>> go mod init
>> kubebuilder init —domain gov.io

```

#### Step 2: Create 3 different API versions for kind Registeration

Now let's create APIs for kind Registration using

```bash
>> Kubebuilder create api —group cowin —kind Registration —version v1
    Create resource(y/n)  - y
    Create controller(y/n)  - y    
```
Similarly we create other API versions also, so finally output will be

```bash
// v1 API
type RegistrationSpec struct {
	Name             string `json:"name,omitempty"`
	VerifiedID       string `json:"verified_id,omitempty"`
	RegistrationDate string `json:"registration_date,omitempty"`
}

// v2 API
type RegistrationSpec struct {
	Name             string `json:"name,omitempty"`
	VerifiedID       string `json:"verified_id,omitempty"`
	RegistrationDate string `json:"registration_date,omitempty"`
	VaccineName      string `json:"vaccine_name,omitempty"`
}

// v3 API
type RegistrationSpec struct {
	Name           string           `json:"name,omitempty"`
	VerifiedID     string           `json:"verified_id,omitempty"`
	VaccineDetails []*VaccineDetail `json:"vaccine_details,omitempty"`
}

type VaccineDetail struct {
	RegistrationDate string `json:"registration_date,omitempty"`
	VaccineName      string `json:"vaccine_name,omitempty"`
}

```
