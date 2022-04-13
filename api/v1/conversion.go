package v1

import (
	"fmt"
	conv "k8s.io/apimachinery/pkg/conversion"
	v3 "kube-vaccine/api/v3"
	"kube-vaccine/util"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

var setupLog = ctrl.Log.WithName("setup")

func (src *Registration) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v3.Registration)

	if err := Convert_v1_Registration_To_v3_Registration(src, dst, nil); err != nil {
		return err
	}

	restored := &v3.Registration{}

	ok, err := util.UnmarshalData(src, restored)
	if err != nil {
		return err
	}

	if ok == true {
		if restored.Spec.VaccineDetails == nil {
			return fmt.Errorf("Vaccine details should not be nil in v3")
		}
		dst.Spec.VaccineDetails = restored.Spec.VaccineDetails
	} else {
		dst.Spec.VaccineDetails = []*v3.VaccineDetail{
			&v3.VaccineDetail{
				VaccineName:      "Covishield",
				RegistrationDate: src.Spec.RegistrationDate,
			},
		}
	}
	return nil
}

func (dst *Registration) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v3.Registration)
	if err := Convert_v3_Registration_To_v1_Registration(src, dst, nil); err != nil {
		return err
	}

	if err := util.MarshalData(src, dst); err != nil {
		return err
	}

	if src.Spec.VaccineDetails == nil || len(src.Spec.VaccineDetails) == 0 {
		return fmt.Errorf("Vaccine details should not be empty in V3")
	}
	dst.Spec.RegistrationDate = src.Spec.VaccineDetails[0].RegistrationDate

	return nil
}

func Convert_v1_RegistrationSpec_To_v3_RegistrationSpec(in *RegistrationSpec, out *v3.RegistrationSpec, s conv.Scope) error {
	return autoConvert_v1_RegistrationSpec_To_v3_RegistrationSpec(in, out, s)
}

func Convert_v3_RegistrationSpec_To_v1_RegistrationSpec(in *v3.RegistrationSpec, out *RegistrationSpec, s conv.Scope) error {
	return autoConvert_v3_RegistrationSpec_To_v1_RegistrationSpec(in, out, s)
}
