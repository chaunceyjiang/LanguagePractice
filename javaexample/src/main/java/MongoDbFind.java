
import com.mongodb.MongoClient;
import com.mongodb.MongoClientURI;
import com.mongodb.client.FindIterable;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoCursor;
import com.mongodb.client.MongoDatabase;
import org.bson.Document;

import java.io.Serializable;
import java.util.*;

public class MongoDbFind {

    public static void main(String[] args) {
        MongoClientURI uri = new MongoClientURI("mongodb://192.168.1.34/admin");
        MongoClient client = new MongoClient(uri);
        MongoDatabase db = client.getDatabase("cbms");
        MongoCollection<Document>  coll = db.getCollection("main_app_datapath");
        FindIterable<Document> doc = coll.find();
        MongoCursor<Document> iter = doc.iterator();

        while (iter.hasNext()){
            Document c = iter.next();
            if(((String)c.get("state")).equals("apply")){
                System.out.println(c);
            }
            System.out.println(c);
            Object areas = c.get("areas");
            List l = (ArrayList<Document>) c.get("intfs");
            for (Object x:l){
                String server_id =(String)((Document)x).get("server_id");
                System.out.println(server_id);
            }
            System.out.println();
//            System.out.println((int)c.get("instance_id"));

        }
    }
}

class MyMap extends AbstractMap implements Cloneable, Serializable {

    @Override
    public Set<Entry> entrySet() {
        return null;
    }

}