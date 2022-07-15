package errx

const OK uint32 = 200

//prefix - 3 => service error
//postfix - 3 => specific error of a service

//Global Error Code

const SERVER_COMMON_ERROR uint32 = 100001
const REQ_PARAM_ERROR uint32 = 100002
const TOKEN_EXPIRED_ERROR uint32 = 100003
const TOKEN_GENERATE_ERROR uint32 = 100004
const TOKEN_INVALID_ERROR uint32 = 100005
const DB_ERROR uint32 = 100006
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100007

//User service - 110
const USER_NOT_EXIST uint32 = 110001
const USER_HAS_BEEN_REGISTERED uint32 = 110002

//Oder Service - 120

//....
