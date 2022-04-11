package v1

import (
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conv "k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	v2 "kube-vaccine/api/v2"
)

var setupLog = ctrl.Log.WithName("setup")

func (src *Registration) ConvertTo(dstRaw conversion.Hub) error {
	setupLog.Info("ConvertTo called from v1-->v2 \n\n ")
	dst := dstRaw.(*v2.Registration)

	if err := Convert_v1_Registration_To_v2_Registration(src, dst, nil); err != nil {
		return err
	}

	restored := &v2.Registration{}
	setupLog.Info("CONVERT TO1", "src", src)
	setupLog.Info("CONVERT TO1", "dst", dst)
	setupLog.Info("CONVERT TO1", "restored", restored)

	ok, err := UnmarshalData(src, restored)
	if err != nil {
		return err
	}

	//setupLog.Info("CONVERT TO", "key", fmt.Sprintf("2 dst.Spec.Name,dst.Spec.VerifiedID,dst.Spec.RegistrationDate,dst.Spec.VaccineName : %+v %+v %+v %+v", dst.Spec.Name, dst.Spec.VerifiedID, dst.Spec.RegistrationDate, dst.Spec.VaccineName))
	//setupLog.Info("CONVERT TO", "key", fmt.Sprintf("2 restored.Spec.Name,restored.Spec.VerifiedID,restored.Spec.RegistrationDate,restored.Spec.VaccineName : %+v %+v %+v %+v", restored.Spec.Name, restored.Spec.VerifiedID, restored.Spec.RegistrationDate, restored.Spec.VaccineName))
	setupLog.Info("CONVERT TO1", "src", src)
	setupLog.Info("CONVERT TO2", "dst ", dst)
	setupLog.Info("CONVERT TO2", "restored", restored)
	//
	//fmt.Printf("2 dst.Spec.Name,dst.Spec.VerifiedID,dst.Spec.RegistrationDate,dst.Spec.VaccineName : %+v %+v %+v %+v", dst.Spec.Name, dst.Spec.VerifiedID, dst.Spec.RegistrationDate, dst.Spec.VaccineName)
	//fmt.Printf("2 restored.Spec.Name,restored.Spec.VerifiedID,restored.Spec.RegistrationDate,restored.Spec.VaccineName : %+v %+v %+v %+v", restored.Spec.Name, restored.Spec.VerifiedID, restored.Spec.RegistrationDate, restored.Spec.VaccineName)

	if ok == true {
		if restored.Spec.VaccineName == "" {
			return fmt.Errorf("Vaccine name cannot be empty in v2")
		}
		dst.Spec.VaccineName = restored.Spec.VaccineName
	} else {
		dst.Spec.VaccineName = "COVISHIELD"
	}

	setupLog.Info("\n\n ConvertTo called Ended v1-->v2 \n\n ")

	return nil
}

func (dst *Registration) ConvertFrom(srcRaw conversion.Hub) error {
	setupLog.Info("ConvertFrom called from v2-->v1 \n\n ")
	src := srcRaw.(*v2.Registration)
	if err := Convert_v2_Registration_To_v1_Registration(src, dst, nil); err != nil {
		return err
	}

	setupLog.Info("ConvertFrom1", "dst", dst)
	setupLog.Info("ConvertFrom1", "src", src)

	if err := MarshalData(src, dst); err != nil {
		return err
	}

	setupLog.Info("ConvertFrom2", "dst", dst)
	setupLog.Info("ConvertFrom2", "src", src)

	setupLog.Info("\n\nConvertFrom ended from v2-->v1 \n")
	return nil
}

func Convert_v2_RegistrationSpec_To_v1_RegistrationSpec(in *v2.RegistrationSpec, out *RegistrationSpec, s conv.Scope) error {
	return autoConvert_v2_RegistrationSpec_To_v1_RegistrationSpec(in, out, s)
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
