package main
import "fmt"
func main(){
	var Arr1[100][100] int
	var Arr2[100][100] int
	var sum[100][100] int
	var row, col int
	fmt.Println("Enter the row1 of array 1")
	fmt.Scanln(&row)
	fmt.Println("Enter the col1 of array 1")
	fmt.Scanln(&col)

	fmt.Println("+++++++++MATRIX1++++++++++++++")

	for i:=0;i<row;i++{
		for j := 0; j <col; j++ {
			fmt.Printf("Enter the element for Matrix1 %d %d :",i+1,j+1)
			fmt.Scanln(&Arr1[i][j])
			
		}
	}
	fmt.Println("+++++++++MATRIX2++++++++++++++")
	for i:=0;i<row;i++{
        for j := 0; j <col; j++ {
			fmt.Printf("Enter the element for Matrix2 %d %d :",i+1,j+1)
			fmt.Scanln(&Arr2[i][j])
			
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			sum[i][j]=Arr1[i][j]+Arr2[i][j]
			
		}
		
	}
	fmt.Print("++++++++++Result+++++++++++")
	fmt.Println()
	
	for i:=0;i<row;i++{
		for j := 0; j <col; j++ {
			fmt.Print(sum[i][j])
			if(j==col-1){
				fmt.Println(" ")
			}

		}
	}
}