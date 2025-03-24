# Horoscope webscraper


> ⚠️ The project is created for educational purposes. It's not used for commercial purposes.


## Route usage

The route: ```http://localhost:3000/horoscope?horoscope={horoscope}&time={time}```


### Queries

horoscope:
- (ARIES) oven
- (TAURUS) telec
- (GEMINI) bliznaci
- (CANCER) rak
- (LEO) luv
- (VIRGO) deva
- (LIBRA) vezni
- (SCORPIO) skorpion
- (SAGITTARIUS) strelec
- (CAPRICORN) kozirog
- (AQUARIUS) vodolei
- (PISCES) ribi

time:
- daily
- weekly
- monthly
- yearly (Default)

### Example Response

```json
{
    "message": "Ако днес имате енергия, използвайте я, за да се освободите от ненужни неща. Може да направите основно почистване или да се разделите с вещи, които вече не ви трябват. Задържането на старото само ще ви пречи да създадете място за нови възможности."
}
```
