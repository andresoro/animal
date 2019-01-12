### Generate unique IDs of the form AdjetiveAdjetiveAnimal

```golang
package main

import "github.com/andresoro/animal"

func main() {
    
    for i := 0; i < 5; i++ {
        fmt.Println(animal.New())
    }
    //MidZoisiteMantaray
    //CorroborantValuedWallaby
    //EthnogenicPlatonicAsianporcupine
    //EnragedShadowgraphicCoelacanth
    //EffervesciblePepperishBuffalo

}
```

This package can generate 28,000,000,000 possible combinations so it can be used as unique url generally safely. I used "math/rand" package methods to generate pseudo random outputs so I can't say there won't be duplicates.

