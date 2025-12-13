package main

import "fmt"

const dividerBar = "============================================="

type Product struct {
	Item     ProductItem
	Quantity int
}

type Store struct {
	Money    int
	Products []*Product
}

// Product 객체의 생성자
func NewProduct(item ProductItem, quantity int) *Product {
	product := new(Product)
	product.Item = item
	product.Quantity = quantity
	return product
}

// Product 상품을 고객에게 서빙
func (p *Product) Serve() error {
	err := p.Item.Make()

	if err != nil {
		return err
	}

	//상품을 포장함
	err = p.Item.Package()

	if err != nil {
		return err
	}

	//상품을 집고 건넴
	err = p.Item.Pick()

	if err != nil {
		return err
	}

	return nil
}

// Store 객체의 생성자
func NewStore(money int, products []*Product) *Store {
	store := new(Store)
	store.Money = money
	store.Products = products
	return store
}

func (s *Store) SellProduct(productName string, quantity int) {
	product := s.GetProduct(productName)

	err := product.Serve()
	if err != nil {
		fmt.Printf("상품을 판매하는 과정 중 문제가 발생했습니다. %v\n", err)
		return
	}

	product.Quantity -= quantity
	s.Money += product.Item.Price() * quantity
}

func (s *Store) GetProducts() []*Product {
	return s.Products
}

func (s *Store) GetProduct(productName string) *Product {
	for _, product := range s.Products {
		if product.Item.Name() == productName {
			return product
		}
	}

	return nil
}

func (s *Store) CheckProductQuantity(
	productName string, quantity int) bool {
	product := s.GetProduct(productName)
	return product.Quantity >= quantity
}

//사용자에게 구매 상품과 수량을 입력받는 함수
func HandleChoiceProduct(myStore *Store) (exit bool) {
	for {
		var choice string
		fmt.Println(dividerBar)
		fmt.Print("구매할 상품의 이름을 알려주세요 (나가기 exit): ")
		fmt.Scanln(&choice)

		if choice == "exit" {
			fmt.Println("이용해주셔서 감사합니다!")
			fmt.Println(dividerBar)
			fmt.Printf("최종 가게 보유 자금 %d원\n", myStore.Money)
			return true
		}

		product := myStore.GetProduct(choice)
		if product == nil {
			fmt.Printf("우리 매장에는 %s 라는 이름의 상품은 없네요.\n", choice)
			continue
		} else if product.Quantity == 0 {
			fmt.Printf("%s의 재고 수량이 떨어졌네요...\n", choice)
			continue
		}

		var quantity int
		fmt.Print("몇 개를 구매하시겠어요?: ")
		fmt.Scanln(&quantity)

		isExists := myStore.CheckProductQuantity(
			product.Item.Name(),
			quantity)
		if isExists == false {
			fmt.Printf("%s의 재고 수량이 떨어졌네요...\n", choice)
			continue
		}

		myStore.SellProduct(product.Item.Name(), quantity)
		fmt.Println("Thank you for using our store!")
		break
	}

	return false
}

func main() {
	americano := NewCoffee("아메리카노", 3000, "블랜딩 커피", Bitter, Waiting)
	latte := NewCoffee("카페라떼", 3500, "블랜딩 커피", Sweet, Waiting)
	cafeMocha := NewCoffee("카페모카", 4000, "디저트 커피", Sweet, Waiting)
	dripCoffee := NewCoffee("드립커피", 7000, "원두 커피", FruitFlavor, Waiting)
	dutchCoffee := NewCoffee("더치커피", 5000, "더치 커피", Bitter, Waiting)

	//판매할 제품 수량을 Product 객체로 생성
	productAmericano := NewProduct(americano, 5)
	productLatte := NewProduct(latte, 2)
	productCafeMocha := NewProduct(cafeMocha, 3)
	productDripCoffee := NewProduct(dripCoffee, 4)
	productDutchCoffee := NewProduct(dutchCoffee, 6)

	myStore := NewStore(0, []*Product{
		productAmericano,
		productLatte,
		productCafeMocha,
		productDripCoffee,
		productDutchCoffee})

	fmt.Printf("store: %v\n", myStore)

	for {
		fmt.Println(dividerBar)
		fmt.Println("welcome to our stroe")
		fmt.Printf("our store's asset is: %d원\n", myStore.Money)
		fmt.Println(dividerBar)
		for i, product := range myStore.GetProducts() {
			fmt.Printf("%d. %s product: %d원, (stuck: %d)\n", i+1,
				product.Item.Name(),
				product.Item.Price(),
				product.Quantity)
		}

		isExit := HandleChoiceProduct(myStore)
		if isExit {
			return
		}
	}
}
