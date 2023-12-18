package main

func main() {}

// O(n)
func MaxProductDifference(nums []int) int {
  w, x, y, z := 0,0,10001,10001
  
  for i := 0; i < len(nums); i++ {
    if nums[i] > w {
      x = w
      w = nums[i]
    } else if nums[i] > x {
      x = nums[i]
    }

    if nums[i] < y {
      z = y
      y = nums[i]
    } else if nums[i] < z {
      z = nums[i]
    } 
  }
  return w * x - y * z
}
