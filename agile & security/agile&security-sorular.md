# Agile & Security Soruları
## Soru 1
Şirket içerisinde geliştirilecek olan bir uygulamanın planlama toplantısına girdin ve toplantıda veri tabanı işlemlerinin yapılacağı modülü senin geliştirmen istendi. Sana geliştireceğin modülde kullanıcı parolalarını veri tabanında `encode`, `hashing`, `encription` yöntemlerinin hangisini kullanarak saklayacağın soruldu. Takım senden bunların ne olduğunu açıklamanı, birbiri arasındaki farkları, hangisi hangi amaçla kullanıldığını açıklamanı bekliyor. Tabi kullanıcı parolalarını saklamak için de doğru yolu belirlemeni istiyor. Bu gereksinimleri karşılayacak cevapları aşağıda bizimle paylaşabilir misin?
## Soru 2
Scrum takımı olarak çalıştığın bir firmada herkesin geliştirdiği bir koda diğer bir takım üyesi code review yapıyor ve eksik noktaları ona ileterek kodun uygun hale gelmesi için birbirinin kodunu inceliyor. Sen de takımın en tecrübeli ve egosu yüksek üyesinin geliştirdiği kodda bir **IDOR(Insecure Direct Object References)** zafiyeti tespit ettin. Bu zafiyetin ne olduğunu ve bu zafiyetin ortadan kaldırılması için nasıl bir çözüm uygulanması gerektiği ile ilgili bu egosu yüksek takım arkadaşına birebirde bilgi vermen gerekiyor. Bu durumda kullanacağın yöntemi, üslup ve bilgi içeriğini aşağıda bizimle paylaşır mısın?
## Soru 3
Agile `Scrum` Metodunda bulunan rolleri işlevleriyle birlikte ve ritüelleri(toplantıları) amaçlarıyla birlikte açıklayınız.
## Soru 4
Git teknolojisi ile ilgili olarak `Repository`, `Branch`, `Commit`, `Pull`, `Push`, `Merge Request`(MR), `Merge Conflict`, `Fork` terimlerinin neler olduğunu açıklayınız.
## Soru 5
Bir git repositorysi için, branch last commit nedir? Repository last commit nedir?
## Soru 6
`SQL` veri tabanları ile `NoSQL` veri tabanları arasındaki farklar nelerdir? Ne zaman hangi teknolojiyi tercih etmeliyiz?
## Soru 7
Basit bir e-ticaret firmasının kullanması için Ürün-Müşteri-Şipariş süreçlerinin gerçekleştirileceği bir modül için oluşturulması istenen veri tabanı modelinin Entity diagramını oluşturunuz. ( [Draw.io](https://draw.io/) sitesini kullanarak şemeyı oluşturup bize iletiniz.)
## Soru 8
```javascript
String GenerateReceiptURL(String baseUrl) {
Random ranGen = new Random();
ranGen.setSeed((new Date()).getTime());
public static final String NFE = "Success"
log.info(NFE);
return (baseUrl + ranGen.nextInt(400000000) + ".html"); }
```
Yukarıdaki kod parçasında hangi zafiyet bulunmaktadır?
<ol type="A">
  <li>Null Dereference</li>
  <li>Privacy Violation</li>
  <li>Log Forging</li>
  <li>Unreleased Resource: Stream</li>
  <li>Insecure Randomless</li>
</ol>

## Soru 9
`Üç etmenli kimlik doğrulama` yapmak için aşağıdakilerden hangilerini seçmeliyiz?
<ol type="A">
  <li>Kullanıcı Adı, Cep telefonu OTP, Parmak İzi</li>
  <li>Kullanıcı Adı, Cep telefonu OTP, Akıllı Kart</li>
  <li>Şifre, Cep telefonu OTP, Akıllı Kart</li>
  <li>Kullanıcı Adı, Şifre, Cep telefonu OTP</li>
  <li>Şifre, Retina, Parmak İzi</li>
</ol>

## Soru 10
```
 public Task<Credentials> GetCustomer(string username, string password)
 {
string query = "SELECT * FROM customers Where
password='"+password+ "' and username='" + username + "';";
DatabaseAccess db = new DatabaseAccess();
DataTable dTable = await db.ExecuteAsync(query, new List<MySqlParameter>
{
```
Yukarıdaki kod satırında `OWASP Top10` zafiyetlerinden biri bulunmaktadır. Hangisinin bulunduğunu ve var olan zafiyetin çözümü için kodun revize edilmiş halini yazınız. ([Kaynağa](https://owasp.org/Top10/) başvurarak araştırma yapabilirsiniz.)

