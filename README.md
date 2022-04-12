# Kube-Vaccine
![visitor badge](https://visitor-badge.glitch.me/badge?page_id=shivi28.kube-vaccine)
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

Let's first create v1 API for kind Registration

```bash
>> Kubebuilder create api —group cowin —kind Registration —version v1
    Create resource(y/n)  - y
    Create controller(y/n)  - y    
```
**Output is:** v1 API created with following schema
```bash
type RegistrationSpec struct {
	Name             string `json:"name,omitempty"`
	VerifiedID       string `json:"verified_id,omitempty"`
	RegistrationDate string `json:"registration_date,omitempty"`
}
```

Now create v2 API

```bash
>> Kubebuilder create api —group cowin —kind Registration —version v2
    Create resource(y/n)  - y
    Create controller(y/n)  - n  
```
**Output:**
```bash
type RegistrationSpec struct {
	Name             string `json:"name,omitempty"`
	VerifiedID       string `json:"verified_id,omitempty"`
	RegistrationDate string `json:"registration_date,omitempty"`
	VaccineName      string `json:"vaccine_name,omitempty"`
}
```

At last create v3 API

```bash
>> Kubebuilder create api —group cowin —kind Registration —version v3
    Create resource(y/n)  - y
    Create controller(y/n)  - n  
```
**Output:**

`v3 API contains`

```bash
...
type RegistrationSpec struct {
	Name           string           `json:"name,omitempty"`
	VerifiedID     string           `json:"verified_id,omitempty"`
	VaccineDetails []*VaccineDetail `json:"vaccine_details,omitempty"`
}

type VaccineDetail struct {
	RegistrationDate string `json:"registration_date,omitempty"`
	VaccineName      string `json:"vaccine_name,omitempty"`
}
...
```
