package cache

//import (
//	"card-validator-service/internal/domain/model"
//	"math/rand"
//	"sync"
//	"testing"
//	"time"
//)
//
//func TestCardIssuerCache_AddAndGet(t *testing.T) {
//	cache := NewCardIssuerCache(10)
//
//	iin := model.IINRange{Start: 4, End: 4}
//	issuer := &model.CardIssuer{
//		IINRanges: issuerName,
//	}
//
//	cache.Add(issuerName, issuer)
//
//	retrievedIssuer := cache.Get(issuerName)
//	if retrievedIssuer == nil {
//		t.Errorf("Expected issuer to be found, but got nil")
//	}
//
//	if retrievedIssuer.Name != issuer.Name {
//		t.Errorf("Expected issuer %+v, but got %+v", issuer, retrievedIssuer)
//	}
//}
//
//func TestCardIssuerCache_GetNonExistentIssuer(t *testing.T) {
//	cache := NewCardIssuerCache(10)
//
//	issuerName := model.IssuerName("MasterCard")
//	retrievedIssuer := cache.Get(issuerName)
//
//	if retrievedIssuer != nil {
//		t.Errorf("Expected nil, but got %+v", retrievedIssuer)
//	}
//}
//
//func TestCardIssuerCache_LRUEviction(t *testing.T) {
//	cache := NewCardIssuerCache(2)
//
//	issuer1 := &model.CardIssuer{Name: "Visa"}
//	issuer2 := &model.CardIssuer{Name: "MasterCard"}
//	issuer3 := &model.CardIssuer{Name: "Amex"}
//
//	cache.Add("Visa", issuer1)
//	cache.Add("MasterCard", issuer2)
//	cache.Add("Amex", issuer3)
//
//	retrievedIssuer := cache.Get("Visa")
//	if retrievedIssuer != nil {
//		t.Errorf("Expected nil for 'Visa', but got %+v", retrievedIssuer)
//	}
//
//	retrievedIssuer = cache.Get("MasterCard")
//	if retrievedIssuer == nil {
//		t.Errorf("Expected issuer 'MasterCard' to be found, but got nil")
//	}
//
//	retrievedIssuer = cache.Get("Amex")
//	if retrievedIssuer == nil {
//		t.Errorf("Expected issuer 'Amex' to be found, but got nil")
//	}
//}
//
//func TestCardIssuerCache_ConcurrentAccess(t *testing.T) {
//	random := rand.New(rand.NewSource(time.Now().UnixNano()))
//	cache := NewCardIssuerCache(10)
//
//	issuers := []model.IssuerName{"Visa", "Master", "Red", "YellowBank"}
//
//	var wg sync.WaitGroup
//	numRoutines := 1000
//
//	for i := 0; i < numRoutines; i++ {
//		wg.Add(1)
//
//		go func(i int) {
//			defer wg.Done()
//
//			index := random.Intn(len(issuers))
//			issuer := issuers[index]
//			cache.Add(issuer, &model.CardIssuer{
//				Name: issuer,
//			})
//
//			retrievedIssuer := cache.Get(issuers[index])
//			if retrievedIssuer == nil {
//				t.Errorf("Routine %d: Expected issuer to be found, but got nil", i)
//			} else if retrievedIssuer.Name != issuer {
//				t.Errorf("Routine %d: Expected issuer %+v, but got %+v", i, issuer, retrievedIssuer)
//			}
//		}(i)
//	}
//
//	wg.Wait()
//}
