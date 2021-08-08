package br.edu.ifpb.gugawag.so.sockets;

import java.io.DataInputStream;
import java.io.DataOutputStream;
import java.io.IOException;
import java.net.ServerSocket;
import java.net.Socket;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;


public class Servidor2 {
    private static final String DIR = "."; // @@@

    public static void main(String[] args) throws IOException {
        System.out.println("== Servidor ==");

        // Configurando o socket
        ServerSocket serverSocket = new ServerSocket(8080);
        Socket socket = serverSocket.accept();

        DataOutputStream sockOut = new DataOutputStream(socket.getOutputStream());
        DataInputStream sockIn = new DataInputStream(socket.getInputStream());

        // laço infinito do servidor
        while (true) {
            System.out.println("Cliente: " + socket.getInetAddress());

            String mensagem = sockIn.readUTF();
            String[] splittedMessage = mensagem.split(" ");
            String cmd = splittedMessage[0]; // this is the command such as create, readdir etc
            try {
                switch(cmd) {
                    case "create": {
                        if (splittedMessage.length >= 2) {
                            String arg = splittedMessage[1];
                            Path p = Paths.get(DIR + "/"+arg);
                            Files.createDirectory(p);
                            sockOut.writeUTF("OK");
                        }
                        break;
                    }
                    case "remove": {
                        if (splittedMessage.length >= 2) {
                            String arg = splittedMessage[1];
                            Path p = Paths.get(DIR + "/"+arg);
                            Files.delete(p);
                            sockOut.writeUTF("OK");
                        }
                        break;
                    }
                    case "rename": {
                        if (splittedMessage.length >= 3) {
                            String arg1 = splittedMessage[1];
                            String arg2 = splittedMessage[2];
                            Path p1 = Paths.get(DIR +"/" + arg1);
                            Path p2 = Paths.get(DIR +"/" + arg2);
                            Files.move(p1, p2);
                            sockOut.writeUTF("OK");
                        }
                        break;
                    }
                    case "readdir": {
                        String subDir = ".";
                        if (splittedMessage.length == 2) {
                            subDir = splittedMessage[1];
                        }
                        String response = "";
                        Path path = Paths.get(DIR+"/"+subDir);
                        Stream<Path> stream = Files.list(path);
                        List<Path> pathList = stream.collect(Collectors.toList());
                        for (Path p: pathList) {
                            response = response.concat(p.getFileName().toString()).concat("\n");
                        }

                        sockOut.writeUTF("\n"+response);
                        break;
                    }
                    case "exit": {
                        sockOut.writeUTF("OK");
                        socket.close();
                        return;
                    }
                    default:
                        sockOut.writeUTF("ERR");
                        break;
                }
            } catch (Exception ex) {
                // some EOF or IO exceptions
                System.out.println(ex.getMessage());
                sockOut.writeUTF("ERR");
            }

        }

    }
}
