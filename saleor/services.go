package saleor

type SaleorManifestService struct {
	SaleorConfigRepository *SaleorConfigRepository
}

func NewSaleorManifestService(saleorConfigRepository *SaleorConfigRepository) *SaleorManifestService {
	return &SaleorManifestService{
		SaleorConfigRepository: saleorConfigRepository,
	}
}

func (s *SaleorManifestService) RegisterSaleorDomain(config *SaleorConfig) (*SaleorConfig, error) {
	return s.SaleorConfigRepository.RegisterSaleorDomain(config)
}

func (s *SaleorManifestService) GetConfigByDomain(domain string) (*SaleorConfig, error) {
	return s.SaleorConfigRepository.GetConfigByDomain(domain)
}
