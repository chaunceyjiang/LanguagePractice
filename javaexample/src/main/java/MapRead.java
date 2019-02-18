package msgPack.msgPack;

import java.io.BufferedInputStream;
import java.io.BufferedReader;
import java.io.ByteArrayOutputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.io.InputStream;
import java.util.HashMap;
import java.util.Map;

import org.msgpack.MessagePack;
import org.msgpack.type.Value;
import org.msgpack.type.ValueType;
import org.msgpack.unpacker.BufferUnpacker;
import org.msgpack.unpacker.Converter;
import org.msgpack.unpacker.Unpacker;

public class MapRead<K, V> {
	private MessagePack msgPack = new MessagePack();
	
	// private BufferUnpacker bufferUnpacker = null;
	private Unpacker bufferUnpacker = null;
	
	public MapRead() {
	}
	
	public void ReadMap(Map<K, V> v, Class<K> keyClass, Class<V> vClass, byte[] data) throws Exception {
		
		InputStream in = new FileInputStream(new File("C:/Users/蒲亮/Desktop/20190102165300_20190102165300_3.btr.pack"));
		this.bufferUnpacker = this.msgPack.createUnpacker(in);
		
		this.bufferUnpacker.resetReadByteCount();
		
		if (this.bufferUnpacker.trySkipNil()) {
			return;
		}
		
		while(true) {
			System.out.println(this.bufferUnpacker.getReadByteCount());
			// Map开始
			int size = this.bufferUnpacker.readMapBegin();
			System.out.println("Map大小:" + size);
			
			for (int i = 0; i < size; i++) {
				Value kValue = this.bufferUnpacker.readValue();
				Value Vvalue = this.bufferUnpacker.readValue();
				System.out.println(kValue + ":" + Vvalue);
			}
			this.bufferUnpacker.readMapEnd();
		}
		
	}
	
	public static byte[] toByteArray(String filename) throws IOException{  
        
        File f = new File(filename);  
        if(!f.exists()){  
            throw new FileNotFoundException(filename);  
        }  
  
        ByteArrayOutputStream bos = new ByteArrayOutputStream((int)f.length());  
        BufferedInputStream in = null;  
        try{  
            in = new BufferedInputStream(new FileInputStream(f));  
            int buf_size = 1024;  
            byte[] buffer = new byte[buf_size];  
            int len = 0;  
            while(-1 != (len = in.read(buffer,0,buf_size))){  
                bos.write(buffer,0,len);  
            }  
            return bos.toByteArray();  
        }catch (IOException e) {  
            e.printStackTrace();  
            throw e;  
        }finally{  
            try{  
                in.close();  
            }catch (IOException e) {  
                e.printStackTrace();  
            }  
            bos.close();  
        }  
    }  
	
	public static void main(String[] args) throws IOException {
		byte[] bytes = toByteArray("C:/Users/蒲亮/Desktop/20190102165300_20190102165300_3.btr.pack");
		
		Map<String, String> map1 = new HashMap<String, String>();
		
		MapRead m = new MapRead<String, String>();
		try {
			m.ReadMap(map1, String.class, String.class, bytes);
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
}
