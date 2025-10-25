package saleor

type SaleorService struct {
	SaleorConfigRepository *SaleorConfigRepository
}

func NewSaleorManifestService(saleorConfigRepository *SaleorConfigRepository) *SaleorService {
	return &SaleorService{
		SaleorConfigRepository: saleorConfigRepository,
	}
}

func (s *SaleorService) RegisterSaleorDomain(config *SaleorConfig) (*SaleorConfig, error) {
	return s.SaleorConfigRepository.create(config)
}

func (s *SaleorService) GetConfigByDomain(domain string) (*SaleorConfig, error) {
	return s.SaleorConfigRepository.getByDomain(domain)
}
