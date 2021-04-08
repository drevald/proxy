import java.net.*;

public class Client {
    public static void main (String[] args) throws Exception {
        System.out.println("Hello Client");
        // SocketAddress addr = new InetSocketAddress("webcache.mydomain.com", 8080);
        // Proxy proxy = new Proxy(Proxy.Type.SOCKS, addr);
        URL url = new URL("http://localhost:8080/server");
        //URLConnection conn = url.openConnection(proxy);
        URLConnection conn = url.openConnection();
        conn.connect();
        Object obj = conn.getContent();
        System.out.println(obj);
    }
}