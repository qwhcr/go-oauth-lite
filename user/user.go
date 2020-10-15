package user

import (
    //"log"

    //"golang.org/x/crypto/bcrypt"
)

//type User struct {
    //ID string
    //Email string
    //Password string
//}


//type StoredUser struct {
    //ID string
    //Email string
    //PasswordHash string
//}

//func PersistNewUser(user *User, 

//func hashAndSalt(pwd string) string {

    //// Use GenerateFromPassword to hash & salt pwd.
    //// MinCost is just an integer constant provided by the bcrypt
    //// package along with DefaultCost & MaxCost.
    //// The cost can be any value you want provided it isn't lower
    //// than the MinCost (4)
    //hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
    //if err != nil {
            //log.Println(err)
    //}
    //// GenerateFromPassword returns a byte slice so we need to
    //// convert the bytes to a string and return it
    //return string(hash)
//}

//func checkPassword (passwordHash string, plainPassword string) bool {
        //// Since we'll be getting the hashed password from the DB it
        //// will be a string so we'll need to convert it to a byte slice
        //byteHash := []byte(passwordHash)
        //err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPassword))
        //if err != nil {
                //return false
        //}

        //return true
//}

