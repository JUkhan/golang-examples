package urlshort

import "testing"

func TestHandler(t *testing.T)  {
	YAMLHandler()
	
}

func BenchmarkXxx(b *testing.B) {
		for i:=0; i< b.N; i+=1{
			YAMLHandler()
		}
}