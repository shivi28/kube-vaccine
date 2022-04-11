package v3

import (
	"encoding/json"
	"fmt"
	ctrl "sigs.k8s.io/controller-runtime"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conv "k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	v2 "kube-vaccine/api/v2"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

var setupLog = ctrl.Log.WithName("setup")

func (src *Registration) ConvertTo(dstRaw conversion.Hub) error {
	setupLog.Info("ConvertTo called from v3-->v2 \n\n ")
	dst := dstRaw.(*v2.Registration)

	if err := Convert_v3_Registration_To_v2_Registration(src, dst, nil); err != nil {
		return err
	}

	setupLog.Info("CONVERT TO1", "src", src)
	setupLog.Info("CONVERT TO1", "dst", dst)

	if err := MarshalData(src, dst); err != nil {
		return err
	}

	// TODO: src.Spec.VaccineDetails should not be nil, add it in a webhook
	if src.Spec.VaccineDetails != nil {
		dst.Spec.VaccineName = src.Spec.VaccineDetails[0].VaccineName
		dst.Spec.RegistrationDate = src.Spec.VaccineDetails[0].RegistrationDate
	}

	setupLog.Info("CONVERT TO1", "src", src)
	setupLog.Info("CONVERT TO1", "dst", dst)

	setupLog.Info("ConvertTo ended from v3-->v2 \n\n ")
	return nil
}

func (dst *Registration) ConvertFrom(srcRaw conversion.Hub) error {
	setupLog.Info("ConvertFrom started from v2-->v3 \n\n ")
	src := srcRaw.(*v2.Registration)
	if err := Convert_v2_Registration_To_v3_Registration(src, dst, nil); err != nil {
		return err
	}

	restored := &Registration{}

	setupLog.Info("ConvertFrom 1", "src", src)
	setupLog.Info("ConvertFrom 1", "dst", dst)
	setupLog.Info("ConvertFrom 1", "restored", restored)

	ok, err := UnmarshalData(src, restored)
	if err != nil {
		return err
	}
	if ok == true {
		if restored.Spec.VaccineDetails == nil {
			return fmt.Errorf("vaccine details should not be nil in v3")
		}
		dst.Spec.VaccineDetails = restored.Spec.VaccineDetails
	} else {
		setupLog.Info("ConvertFrom 3", "src", src)
		setupLog.Info("ConvertFrom 3", "dst", dst)
		setupLog.Info("ConvertFrom 3", "restored", restored)
		dst.Spec.VaccineDetails = make([]*VaccineDetail, 0)
		dst.Spec.VaccineDetails = append(dst.Spec.VaccineDetails, &VaccineDetail{
			VaccineName:      src.Spec.VaccineName,
			RegistrationDate: src.Spec.RegistrationDate,
		})
	}

	setupLog.Info("ConvertFrom 2", "src", src)
	setupLog.Info("ConvertFrom 2", "dst", dst)
	setupLog.Info("ConvertFrom 2", "restored", restored)

	setupLog.Info("\n\n ConvertFrom ended from v2-->v3 \n\n ")
	return nil
}

func Convert_v3_RegistrationSpec_To_v2_RegistrationSpec(in *RegistrationSpec, out *v2.RegistrationSpec, s conv.Scope) error {
	return autoConvert_v3_RegistrationSpec_To_v2_RegistrationSpec(in, out, s)
}

func Convert_v2_RegistrationSpec_To_v3_RegistrationSpec(in *v2.RegistrationSpec, out *RegistrationSpec, s conv.Scope) error {
	return autoConvert_v2_RegistrationSpec_To_v3_RegistrationSpec(in, out, s)
}

const DataAnnotation = "cowin.gov.in/conversion-data"

// MarshalData stores the source object as json data in the destination object annotations map.
// It ignores the metadata of the source object.
func MarshalData(src metav1.Object, dst metav1.Object) error {
	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(src)
	if err != nil {
		return err
	}
	delete(u, "metadata")

	data, err := json.Marshal(u)
	if err != nil {
		return err
	}
	annotations := dst.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}
	annotations[DataAnnotation] = string(data)
	dst.SetAnnotations(annotations)
	return nil
}

// UnmarshalData tries to retrieve the data from the annotation and unmarshals it into the object passed as input.
func UnmarshalData(from metav1.Object, to interface{}) (bool, error) {
	annotations := from.GetAnnotations()
	data, ok := annotations[DataAnnotation]
	if !ok {
		return false, nil
	}
	if err := json.Unmarshal([]byte(data), to); err != nil {
		return false, err
	}
	delete(annotations, DataAnnotation)
	from.SetAnnotations(annotations)
	return true, nil
}
