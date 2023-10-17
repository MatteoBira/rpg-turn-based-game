package main

import (
	"fmt"
  "math/rand"
)

func main() {
  var player1name string
  var p1fire int 
  var p2fire int
  var p1plz int 
  var p2plz int
	fmt.Println("What's your name: ")
  _, err := fmt.Scanln(&player1name)
  if err != nil {
    fmt.Println("Error:", err)
  }
  player1hp, player2hp := 100, 100
  player2name := "computer"
  for player1hp > 0 && player2hp > 0 {
      var p1atk int
      p2atk := rand.Intn(4)+1
      atkRef := map[int]string{
          1: "Normal Attack",
          2: "Fire Attack",
          3: "Elettric Attack",
          4: "Heal",
      }
      fmt.Println("your HP's:", player1hp)
      fmt.Println("computer HP's:", player2hp)


      if p1plz != 0 {
        fmt.Println("you are paralized, you cannot attack\n")
        p1plz -=1
      } else {
        fmt.Println(player1name, ": what is your move? \n1.Normal Attack (25 dmg) 2.Fire Attack(15 dmg,3 dmgx3urns)\n3.Elettric Attack(20,paralizedX1turn ) 4.heal(+25HP)  ")
        _, err := fmt.Scanln(&p1atk)
        if err != nil {
          fmt.Println("Error:", err)
        }

        switch p1atk {
        case 1:
          player2hp -= 25
        case 2:
          player2hp -= 15
          if rand.Intn(2) == 1 {
            p1fire += 3
          } 
          
        case 3:
          player2hp -= 20
          if rand.Intn(3) == 1 {
            p2plz += 1
          }
        case 4:
          if rand.Intn(2) == 1 {
            player1hp += 25
          }
        }
        fmt.Println("you played", atkRef[p1atk], "\n")
      }

    // burnt damage
    if p1fire != 0 {
      player2hp -= 3
      p1fire -= 1
    
    }
    
    if p2plz != 0 {
      fmt.Println(player2name, "is paralized, he cannot attack")
      p2plz -=1
    } else {
      
      switch p2atk {
        case 1:
          player1hp -= 25
        case 2:
          player1hp -= 15
          if rand.Intn(3) == 1 {
            p1fire += 3
          }
        case 3:
          player1hp -= 20
          if rand.Intn(2) == 1 {
            p1plz += 1
          }
        case 4:
          if rand.Intn(2) == 1 {
            player2hp += 25
          }
      }
      fmt.Println("computer played", atkRef[p2atk], "\n")
    }

    // burnt damage
    if p2fire != 0 {
      player1hp -= 3
      p1fire -= 1
    }
      
    }
  if player2hp < 1 {
    fmt.Println(player1name, " you won🏆!")
  } else {
    fmt.Println(player1name, " you lose💀!")
  }
}
