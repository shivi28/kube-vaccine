package v2

import (
	"fmt"
	v3 "kube-vaccine/api/v3"
	ctrl "sigs.k8s.io/controller-runtime"

	conv "k8s.io/apimachinery/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

var setupLog = ctrl.Log.WithName("setup")

func (src *Registration) ConvertTo(dstRaw conversion.Hub) error {
	setupLog.Info("ConvertTo called from v2-->v3 \n\n ")
	dst := dstRaw.(*v3.Registration)

	if err := Convert_v2_Registration_To_v3_Registration(src, dst, nil); err != nil {
		return err
	}

	setupLog.Info("CONVERT TO1", "src", src)
	setupLog.Info("CONVERT TO1", "dst", dst)

	restored := &v3.Registration{}
	ok, err := UnmarshalData(src, restored)
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
				VaccineName:      src.Spec.VaccineName,
				RegistrationDate: src.Spec.RegistrationDate,
			},
		}
	}

	setupLog.Info("CONVERT TO1", "src", src)
	setupLog.Info("CONVERT TO1", "dst", dst)

	setupLog.Info("ConvertTo ended from v2-->v3 \n\n ")
	return nil
}

func (dst *Registration) ConvertFrom(srcRaw conversion.Hub) error {
	setupLog.Info("ConvertFrom started from v3-->v2 \n\n ")
	src := srcRaw.(*v3.Registration)
	if err := Convert_v3_Registration_To_v2_Registration(src, dst, nil); err != nil {
		return err
	}

	if err := MarshalData(src, dst); err != nil {
		return err
	}

	if src.Spec.VaccineDetails == nil || len(src.Spec.VaccineDetails) == 0 {
		return fmt.Errorf("Vaccine details should not be empty in V3")
	}
	dst.Spec.RegistrationDate = src.Spec.VaccineDetails[0].RegistrationDate
	dst.Spec.VaccineName = src.Spec.VaccineDetails[0].VaccineName

	setupLog.Info("\n\n ConvertFrom ended from v3-->v2 \n\n ")
	return nil
}

func Convert_v3_RegistrationSpec_To_v2_RegistrationSpec(in *v3.RegistrationSpec, out *RegistrationSpec, s conv.Scope) error {
	return autoConvert_v3_RegistrationSpec_To_v2_RegistrationSpec(in, out, s)
}

func Convert_v2_RegistrationSpec_To_v3_RegistrationSpec(in *RegistrationSpec, out *v3.RegistrationSpec, s conv.Scope) error {
	return autoConvert_v2_RegistrationSpec_To_v3_RegistrationSpec(in, out, s)
}
