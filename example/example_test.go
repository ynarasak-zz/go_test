package example
import "testing"
func TestExampleSuccess(t *testing.T){
  result := Example("success")
  if result != 0 {
    t.Fatal("failed test")
  }
}
