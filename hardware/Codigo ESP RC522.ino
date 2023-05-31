#include <Arduino.h>
#include <MFRC522.h> //biblioteca responsável pela comunicação com o módulo RFID-RC522
#include <SPI.h> //biblioteca para comunicação do barramento SPI
#include <HTTPClient.h>
#include <WiFi.h>

HTTPClient http;

#define SS_PIN    21
#define RST_PIN   22
#define SIZE_BUFFER     18
#define MAX_SIZE_BLOCK  16
#define LED_R 14
#define LED_G 13
#define LED_Y 12
#define BIP 27


// Informe as credencias da sua rede wifi
const char* ssid = "Costelini";
const char* password = "Debora@280375";

//esse objeto 'chave' é utilizado para autenticação
MFRC522::MIFARE_Key key;
//código de status de retorno da autenticação
MFRC522::StatusCode status;

// Definicoes pino modulo RC522
MFRC522 mfrc522(SS_PIN, RST_PIN); 

int LED_BUILTIN = 2;

void setup() {
    pinMode(LED_R, OUTPUT);
    pinMode(LED_G, OUTPUT);
    pinMode(LED_Y, OUTPUT);
    pinMode(BIP, OUTPUT);
    Serial.begin(115200);
    delay(3000);
  
    WiFi.begin(ssid, password);
     Serial.println("Conectando no wifi...");
    
  // Aguardando para conectar no wifi
    uint32_t notConnectedCounter = 0;
    while (WiFi.status() != WL_CONNECTED) {
          digitalWrite(LED_Y, HIGH);
          delay(500);
          digitalWrite(LED_Y, LOW);
          delay(500);
        Serial.println("Conectando Wifi...");
        notConnectedCounter++;
        if(notConnectedCounter > 50) { // Após 5 segundos nossa placa é resetada
            Serial.println("Reiniciando placa...");
            ESP.restart();
        }
    }
    
    Serial.println("Wifi conectado, IP address: ");
    Serial.println(WiFi.localIP()); 
    
    SPI.begin(); // Init SPI bus
    pinMode (LED_BUILTIN, OUTPUT);
     
    digitalWrite(BIP, HIGH);
    delay(50);
    digitalWrite(BIP, LOW);
    delay(200);
    digitalWrite(BIP, HIGH);
    delay(50);
    digitalWrite(BIP, LOW);

    Serial.println("Aproxime o seu cartao do leitor..."); 
    
    // Inicia MFRC522
    mfrc522.PCD_Init(); 
    // Mensagens iniciais no serial monitor TESTE  PARA SABER SE ESTA TUDO FUNCNIONANDO
    
    Serial.println();
}
void loop() {
    // Aguarda a aproximacao do cartao
    if ( ! mfrc522.PICC_IsNewCardPresent()) 
    {
      return;
    }
    // Seleciona um dos cartoes
    if ( ! mfrc522.PICC_ReadCardSerial()) 
    {
      digitalWrite(LED_Y, HIGH);
      digitalWrite(LED_G, HIGH);
      digitalWrite(LED_R, HIGH);
      delay(500);
      digitalWrite(LED_Y, LOW);
      digitalWrite(LED_G, LOW);
      digitalWrite(LED_R, LOW);
      delay(500);
      Serial.println("Falha na conexão com o leitor RHIF");
      return;
    }
    Serial.println("Aproxime o seu cartao do leitor..."); 
  
    String conteudo= "";
    byte letra;
    for (byte i = 0; i < mfrc522.uid.size; i++) 
    {
       Serial.println(mfrc522.uid.uidByte[i] < 0x10 ? " 0" : " ");
       Serial.print("Binario: ");
       Serial.println(mfrc522.uid.uidByte[i], BIN);
       conteudo.concat(String(mfrc522.uid.uidByte[i] < 0x10 ? " 0" : " "));
       conteudo.concat(String(mfrc522.uid.uidByte[i], BIN));
    }
    
    Serial.println("Tag : " + String(conteudo));

    Serial.println("Call Api: http://192.168.0.9:9015/attendance");
    http.begin("http://192.168.0.9:9015/attendance");
    http.addHeader("Content-Type", "application/json");
    String httpRequestData = "{\"tag\":\"" + String(conteudo)  + "\"}";
    int httpResponseCode = http.POST(httpRequestData);


      if (httpResponseCode== 200) 
//Double BIP lida cartao
      digitalWrite(BIP, HIGH);
      delay(50);
      digitalWrite(BIP, LOW);
      delay(200);
      digitalWrite(BIP, HIGH);
      delay(50);
      digitalWrite(BIP, LOW);
//Double green lida cartao
      digitalWrite(LED_G, HIGH);
      delay(200);
      digitalWrite(LED_G, LOW);
      delay(200);
      digitalWrite(LED_G, HIGH);
      delay(200);
      digitalWrite(LED_G, LOW);
      delay(200);

      if (httpResponseCode>0) {
      Serial.print("HTTP Response code: ");
      Serial.println(httpResponseCode);
      String payload = http.getString();
      Serial.println(payload);
    }
    else {
      digitalWrite(LED_R, HIGH);
      digitalWrite(BIP, HIGH);
      delay(1000);
      digitalWrite(BIP, LOW);
      digitalWrite(LED_R, LOW);
      Serial.print("Error code: ");
      Serial.println(httpResponseCode);
    }
    // instrui o PICC quando no estado ACTIVE a ir para um estado de "parada"
    mfrc522.PICC_HaltA(); 
    // "stop" a encriptação do PCD, deve ser chamado após a comunicação com autenticação, caso contrário novas comunicações não poderão ser iniciadas
    mfrc522.PCD_StopCrypto1();
}