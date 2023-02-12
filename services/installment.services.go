package services

import (
	"test-kr-sigma/models/entities"
	"test-kr-sigma/models/web"
	"test-kr-sigma/repositories"
)

type InstallmentService interface {
	Create(loanLimitID uint, goodSlug string, userID uint, selectMonth string) (web.InstallmentResponseDTO, error)
}

type InstallmentServiceImpl struct {
	installmentRepository repositories.InstallmentRepository
	userRepository        repositories.UserRepository
	goodRepository        repositories.GoodRepository
}

func NewInstallmentService(InstallmentRepository repositories.InstallmentRepository, UserRepository repositories.UserRepository, GoodRepository repositories.GoodRepository) InstallmentService {
	return &InstallmentServiceImpl{installmentRepository: InstallmentRepository, userRepository: UserRepository, goodRepository: GoodRepository}
}

func ResponseBodyInstallment(installment entities.Installment) web.InstallmentResponseDTO {
	return web.InstallmentResponseDTO{
		OTR:               installment.OTR,
		AdminFee:          installment.AdminFee,
		AmountInstallment: installment.AmountInstallment,
		AmountRate:        installment.AmountRate,
		GoodName:          installment.GoodName,
		UserID:            installment.UserID,
		GoodID:            installment.GoodID,
		LoanLimitID:       installment.LoanLimitID,
	}
}

func (installmentService *InstallmentServiceImpl) Create(loanLimitID uint, goodSlug string, userID uint, selectMonth string) (web.InstallmentResponseDTO, error) {
	user, err := installmentService.userRepository.FindByID(userID)
	loan, err := installmentService.userRepository.FindLoanByID(loanLimitID)
	good, err := installmentService.goodRepository.FindBySlug(goodSlug)

	installment := entities.Installment{
		OTR:               good.Price,
		AdminFee:          150000,
		AmountInstallment: loan.FourthMonth,
		AmountRate:        good.Rate,
		GoodName:          good.Name,
		UserID:            user.ID,
		GoodID:            good.ID,
		LoanLimitID:       loan.ID,
	}

	switch len(selectMonth) > 0 {
	case selectMonth == "fourth_month":
		installment.AmountInstallment = loan.FourthMonth
	case selectMonth == "third_month":
		installment.AmountInstallment = loan.ThirdMonth
	case selectMonth == "second_month":
		installment.AmountInstallment = loan.SecondMonth
	case selectMonth == "first_month":
		installment.AmountInstallment = loan.FirstMonth
	default:
		installment.AmountInstallment = loan.FourthMonth
	}

	createInstallment, err := installmentService.installmentRepository.Create(installment)
	respInstallment := ResponseBodyInstallment(*createInstallment)
	return respInstallment, err
}
