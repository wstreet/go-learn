package src
import(
	"fmt"
)


type SalaryCalculator interface {
	DispalaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type Employee struct {
	firstName string
	lastName string
	age int
}


func (e Employee) DispalaySalary()  {
	
}

func main() {
	
}