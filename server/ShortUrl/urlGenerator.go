package ShortUrl


import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"UrlShortener/server/model"

)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"


/* Retrieve the short URL

@param longUrl string
@param db *sql.DB
@param count integer(uint64)    ---no of entries in the table
@param custom string  			---Custom url by user(default value = default)


@return string  ----shortUrl

*/


func GetShortUrl(longUrl string,custom string)string {

	var shortU string
	var shortUrl string
	hashedValue := GetMD5Hash(longUrl)


	if(custom == "default"){

		shortUrl = "localhost:8882/reshmi/" + Encode(model.Count + 1)

	} else {

		fmt.Println(custom)
		shortUrl = "localhost:8882/reshmi/" + custom

	}

	err := model.Create(hashedValue,longUrl,shortUrl)

	if(err != nil){

		log.Println(err)
		fmt.Println("Error")

		shortU = model.ShowShort(hashedValue)

		return shortU
	}
	return shortUrl

}


/* Hash the long url

@param text string --- long url

@return string  ---hashed string

*/

func GetMD5Hash(text string) string {

	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}


/*Encode the id of teh entry

@param n integer(uint64) --id in the table

@return string --encoded string

*/


func Encode(n uint64) string {

	t := make([]byte, 0)

	/* Special case */
	if n == 0 {
		return string(alphabet[0])
	}

	/* Map */
	for n > 0 {
		r := n % uint64(len(alphabet))
		t = append(t, alphabet[r])
		n = n / uint64(len(alphabet))
	}

	/* Reverse */
	for i, j := 0, len(t) - 1; i < j; i, j = i + 1, j - 1 {
		t[i], t[j] = t[j], t[i]
	}

	fmt.Println(string(t))
	return string(t)
}