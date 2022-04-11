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
	setupLog.Info("ConvertTo called from v1-->v3 \n\n ")
	dst := dstRaw.(*v3.Registration)

	if err := Convert_v1_Registration_To_v3_Registration(src, dst, nil); err != nil {
		return err
	}

	restored := &v3.Registration{}
	setupLog.Info("CONVERT TO1", "src", src)
	setupLog.Info("CONVERT TO1", "dst", dst)
	setupLog.Info("CONVERT TO1", "restored", restored)

	ok, err := util.UnmarshalData(src, restored)
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

	setupLog.Info("\n\n ConvertTo called Ended v1-->v3 \n\n ")

	return nil
}

func (dst *Registration) ConvertFrom(srcRaw conversion.Hub) error {
	setupLog.Info("ConvertFrom called from v3-->v1 \n\n ")
	src := srcRaw.(*v3.Registration)
	if err := Convert_v3_Registration_To_v1_Registration(src, dst, nil); err != nil {
		return err
	}

	setupLog.Info("ConvertFrom1", "dst", dst)
	setupLog.Info("ConvertFrom1", "src", src)

	if err := util.MarshalData(src, dst); err != nil {
		return err
	}

	if src.Spec.VaccineDetails == nil || len(src.Spec.VaccineDetails) == 0 {
		return fmt.Errorf("Vaccine details should not be empty in V3")
	}
	dst.Spec.RegistrationDate = src.Spec.VaccineDetails[0].RegistrationDate

	setupLog.Info("ConvertFrom2", "dst", dst)
	setupLog.Info("ConvertFrom2", "src", src)

	setupLog.Info("\n\nConvertFrom ended from v3-->v1 \n")
	return nil
}

func Convert_v1_RegistrationSpec_To_v3_RegistrationSpec(in *RegistrationSpec, out *v3.RegistrationSpec, s conv.Scope) error {
	return autoConvert_v1_RegistrationSpec_To_v3_RegistrationSpec(in, out, s)
}

func Convert_v3_RegistrationSpec_To_v1_RegistrationSpec(in *v3.RegistrationSpec, out *RegistrationSpec, s conv.Scope) error {
	return autoConvert_v3_RegistrationSpec_To_v1_RegistrationSpec(in, out, s)
}
