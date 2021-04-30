// ─── WRIITEN AND IDEALIZED BY THITI MAHAWANNAKIT ────────────────────────────────
/*
   This can be also designed to have the DI pattern or any other pattern as u desire
   I created this just for the sake of my comfortable and trying to modulize the logics to be reuseable
*/
// ────────────────────────────────────────────────────────────────────────────────

package validator

import (
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator_core "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)


var validatorEntity *validator_core.Validate = validator_core.New()

var translator = en.New()
var uni = ut.New(translator, translator)
var trans ut.Translator

type dynamicStruct = map[string]interface{}

// InitializeTranslator -> A setting up for translator
func InitializeTranslator(){
	transS, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}
	trans = transS

	// Registering Default Translation
	if err := en_translations.RegisterDefaultTranslations(validatorEntity, trans); err != nil {
		log.Fatal(err)
	}
	// ────────────────────────────────────────────────────────────────────────────────


	_ = validatorEntity.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator_core.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = validatorEntity.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator_core.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})
}

// ValidateStructAndGetErrorMsg -> uses for validating any struct and returns if there is any error or not according to the tags rule
func ValidateStructAndGetErrorMsg(data interface{}) (bool,dynamicStruct){
	err := validatorEntity.Struct(data)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	if err != nil {

		validateData := make(dynamicStruct)

		for _, e := range err.(validator_core.ValidationErrors) {
			validateData[e.Field()] = e.Translate(trans)
		}
		return false , validateData
	}else{
		return true , nil
	}
}