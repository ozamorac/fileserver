# fileserver

Servidor que permite transferir archivos entre 2 o más clientes usando un
custom protocol (protocolo no estándar) basado en TCP.
Hecho en GO

Abrir con VSCode

![image](https://user-images.githubusercontent.com/30761344/193153883-ee7ae5de-0d5c-40a1-a9ec-af844d702bb5.png)

# Iniciar servidor 
Escribir en el terminal
go run servidor.go

![image](https://user-images.githubusercontent.com/30761344/193153989-eb4079fb-a888-4403-bf27-f996c5409c87.png)

# Iniciar Cliente (Modo receive)
Escribir en el terminal
go run cliente.go receive 1 cliente1
Donde: 1 es el canal
       cliente1 es el cliente activo en el canal
       
![image](https://user-images.githubusercontent.com/30761344/193154361-a1bf2ec9-214b-425d-8e6b-cee9c0b92da1.png)

# Enviar archivo a canal
Escribir en el terminal
go run cliente.go send 1 file1.txt

![image](https://user-images.githubusercontent.com/30761344/193154515-8e5b6d12-2473-459a-bc63-f23862dae9ad.png)

![image](https://user-images.githubusercontent.com/30761344/193154541-48a964f3-2c19-441f-9ba9-6887e00806eb.png)



