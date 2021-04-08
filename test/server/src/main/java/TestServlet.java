import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.OutputStream;

public class TestServlet extends HttpServlet {

    @Override
    public void service(HttpServletRequest request, HttpServletResponse response) {
        String remoteHost = request.getRemoteHost();
        String remoteAddr = request.getRemoteAddr();
        System.out.println("Remote host is " + remoteHost);
        System.out.println("Remote address is " + remoteAddr);
        try {
            OutputStream os = response.getOutputStream();
            os.write(remoteHost.getBytes());
            os.flush();
            os.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

}

