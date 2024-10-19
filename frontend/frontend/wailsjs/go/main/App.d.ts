// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {user} from '../models';
import {apiresps} from '../models';
import {googleuser} from '../models';

export function GetFriends(arg1:user.User,arg2:string):Promise<apiresps.Friends>;

export function GetSession():Promise<string>;

export function GetUser(arg1:googleuser.GoogleUser):Promise<user.User>;

export function ServerOnline():Promise<boolean>;

export function SignIn():Promise<user.User>;

export function ValidateUsername(arg1:user.User,arg2:string):Promise<apiresps.ValidateUsernameData>;
