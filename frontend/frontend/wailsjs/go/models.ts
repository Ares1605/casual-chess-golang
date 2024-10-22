export namespace apiresps {
	
	export class AwaitSignIn {
	    success: boolean;
	    message?: string;
	    // Go type: struct { Message string "json:\"message\""; Type string "json:\"type\"" }
	    error?: any;
	    // Go type: struct { Token string "json:\"token\""; User user
	    data?: any;
	
	    static createFrom(source: any = {}) {
	        return new AwaitSignIn(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.message = source["message"];
	        this.error = this.convertValues(source["error"], Object);
	        this.data = this.convertValues(source["data"], Object);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Friends {
	    success: boolean;
	    message?: string;
	    // Go type: struct { Message string "json:\"message\""; Type string "json:\"type\"" }
	    error?: any;
	    data?: models.BasicUser[];
	
	    static createFrom(source: any = {}) {
	        return new Friends(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.message = source["message"];
	        this.error = this.convertValues(source["error"], Object);
	        this.data = this.convertValues(source["data"], models.BasicUser);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class JunkResp {
	    success: boolean;
	    message?: string;
	    // Go type: struct { Message string "json:\"message\""; Type string "json:\"type\"" }
	    error?: any;
	    data?: any;
	
	    static createFrom(source: any = {}) {
	        return new JunkResp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.message = source["message"];
	        this.error = this.convertValues(source["error"], Object);
	        this.data = source["data"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class SetupUserData {
	    username: string;
	
	    static createFrom(source: any = {}) {
	        return new SetupUserData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.username = source["username"];
	    }
	}
	export class SetupUser {
	    success: boolean;
	    message?: string;
	    // Go type: struct { Message string "json:\"message\""; Type string "json:\"type\"" }
	    error?: any;
	    // Go type: SetupUserData
	    data?: any;
	
	    static createFrom(source: any = {}) {
	        return new SetupUser(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.message = source["message"];
	        this.error = this.convertValues(source["error"], Object);
	        this.data = this.convertValues(source["data"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ValidateUsernameData {
	    valid: boolean;
	    reason?: string;
	
	    static createFrom(source: any = {}) {
	        return new ValidateUsernameData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.valid = source["valid"];
	        this.reason = source["reason"];
	    }
	}
	export class ValidateUsername {
	    success: boolean;
	    message?: string;
	    // Go type: struct { Message string "json:\"message\""; Type string "json:\"type\"" }
	    error?: any;
	    // Go type: ValidateUsernameData
	    data?: any;
	
	    static createFrom(source: any = {}) {
	        return new ValidateUsername(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.message = source["message"];
	        this.error = this.convertValues(source["error"], Object);
	        this.data = this.convertValues(source["data"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace googlejwt {
	
	export class GoogleJWT {
	    iss: string;
	    azp: string;
	    aud: string;
	    sub: string;
	    email: string;
	    email_verified: boolean;
	    at_hash: string;
	    name: string;
	    picture: string;
	    given_name: string;
	    family_name: string;
	    iat: number;
	    exp: number;
	
	    static createFrom(source: any = {}) {
	        return new GoogleJWT(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.iss = source["iss"];
	        this.azp = source["azp"];
	        this.aud = source["aud"];
	        this.sub = source["sub"];
	        this.email = source["email"];
	        this.email_verified = source["email_verified"];
	        this.at_hash = source["at_hash"];
	        this.name = source["name"];
	        this.picture = source["picture"];
	        this.given_name = source["given_name"];
	        this.family_name = source["family_name"];
	        this.iat = source["iat"];
	        this.exp = source["exp"];
	    }
	}

}

export namespace user {
	
	export class User {
	    id: number;
	    username: string;
	    setup_complete: boolean;
	    google_id: string;
	    profile_url: string;
	    email: string;
	    name: string;
	    decoded_jwt?: googlejwt.GoogleJWT;
	    encoded_jwt: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.username = source["username"];
	        this.setup_complete = source["setup_complete"];
	        this.google_id = source["google_id"];
	        this.profile_url = source["profile_url"];
	        this.email = source["email"];
	        this.name = source["name"];
	        this.decoded_jwt = this.convertValues(source["decoded_jwt"], googlejwt.GoogleJWT);
	        this.encoded_jwt = source["encoded_jwt"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

