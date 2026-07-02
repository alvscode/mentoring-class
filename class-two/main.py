def main():
  
    frutas = ["banana", "maçã", "laranja", "uva", "abacaxi", "melancia", "pera", "manga", "kiwi", "morango"]
        
    # for i in range(0, len(frutas)):
        
    #     if frutas[i] == "abacaxi":
    #         continue   
        
    #     if frutas[i] == "pera":
    #         break 
  
    #     print(frutas[i])
        
        
    quantidadeDeFrutas = len(frutas)
    tentativas = 0   
   
    # while TRUE:
    #  print("Olá, mundo!")
     
    #  tentativas += 1
     
    #  if tentativas == 150:
    #   frutas.append("coco")
    #   quantidadeDeFrutas = len(frutas)
      
      
    nomedocliente =   nome(1)
    print(nomedocliente)
   

def nome(nome: str) -> str:
     return f"Meu nome é {nome}"

main()