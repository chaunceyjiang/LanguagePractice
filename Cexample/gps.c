#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <signal.h>

int main(void)
{	int ds_fd;
	unsigned char ds_buf[2];
	float result_te;
	float latitude;
	float longitude;
	unsigned char ch,status,tmp;

while(1)
{	ch = buf[5];
	status = buf[GetComma(2,buf)];

	if ((ds_fd=open("/dev/DS18B20", O_RDWR|O_NDELAY|O_NOCTTY)) < 0)
	{
		perror("open");
		exit(1);
	}
	if(ch == 'C')
	{
		if (status == 'A')
		{
		latitude = Get_Double_Number(&buf[GetComma(3,buf)]);
		longitude = Get_Double_Number(&buf[GetComma(5,buf)]);
		return 1;
		}
	}
	read(ds_fd, latitude,longitude, 2);
	printf("%.2f,%.2f");
	latitude = ((int)latitude+0.005) / 100;
	ongitude= ((int)longitude +0.005)/ 100;
		{
			FILE *fp;
			fp=fopen("data.txt","w");
			gcvt(longitude,latitude);
			fputs(latitudee,longitude,fp);
			fclose(fp);
		}
		//printf("t = %.1f\n", result_te);
		sleep(20);
	}
	return 0;
}

// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
// #include <sys/types.h>
// #include <sys/stat.h>
// #include <fcntl.h>
// #include <unistd.h>
// #include <signal.h>

// int main(void)
// {
//     int fd;
//     char *com = "/dev/ttyS0";
//     if ((fd = open(com, O_RDWR | O_CREAT, 0777)) < 0)
//     {
//     }
//     else
//     {
//         while (1)

//         {
            
//     }
//     return 0;
// }
// void pms5003ReceiveDataDepare(uint8_t data)
// {
//     static uint8_t state;
//     static uint8_t data_len = 30, data_cnt;
//     static uint8_t rxBuffer[33];

//     if (state == 0 && data == 0x42)
//     {
//         state = 1;
//         rxBuffer[0] = data;
//     }
//     else if (state == 1 && data == 0x4d)
//     {
//         state = 2;
//         rxBuffer[1] = data;
//     }
//     else if (state == 2 && data_len > 0)
//     {
//         data_len--;
//         rxBuffer[2 + data_cnt++] = data;
//         if (data_len == 0)
//             state = 3;
//     }
//     else if (state == 3)
//     {
//         state = 0;
//         data_cnt = 0;
//         data_len = 30;
//         pms5003ReceiveDataAnl(rxBuffer);
//     }
// }
// void pms5003ReceiveDataAnl(uint8_t * data_buffer)
// {
//     uint8_t i;
//     uint32_t sum = 0;

//     for (i = 0; i < 30; i++)
//     {
//         sum += data_buffer[i];
//     }

//     if (sum != ((data_buffer[30] << 8) | data_buffer[31]))
//     {
//         return;
//     }

//     pm25_val = (data_buffer[13] << 8) | data_buffer[14];

//     read(ds_fd, pm25, 2);

//     printf("%n");
//     pm25 = ((int)pm25);
//     {
//         FILE *fp;
//         fp = fopen("data.txt", "w");
//         gcvt(pm25);
//         fputs(pm25, fp);
//         fclose(fp);
//     }

//     //printf("t = %.1f\n", result_te);

//     sleep(20);
// }

// return 0;
// }