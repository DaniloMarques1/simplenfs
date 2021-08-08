package br.edu.ifpb.gugawag.so.sockets;

import java.io.DataInputStream;
import java.io.DataOutputStream;
import java.io.IOException;
import java.net.Socket;
import java.util.Scanner;

public class Cliente2 {

    public static void main(String[] args) throws IOException {
        System.out.println("== Cliente ==");

        // configurando o socket
        Socket socket = new Socket("127.0.0.1", 8080);
        // pegando uma referência do canal de saída do socket. Ao escrever nesse canal, está se enviando dados para o
        // servidor
        DataOutputStream sockOut = new DataOutputStream(socket.getOutputStream());
        // pegando uma referência do canal de entrada do socket. Ao ler deste canal, está se recebendo os dados
        // enviados pelo servidor
        DataInputStream sockIn = new DataInputStream(socket.getInputStream());

        // laço infinito do cliente
        while (true) {
            System.out.print("> ");
            Scanner teclado = new Scanner(System.in);
            String command = teclado.nextLine();
            System.out.println(command);
            // escrevendo para o servidor
            sockOut.writeUTF(command);

            // lendo o que o servidor enviou
            String mensagem = sockIn.readUTF();
            System.out.println("Servidor falou: " + mensagem);
        }
    }
}
