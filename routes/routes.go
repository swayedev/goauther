package routes

// token route method is POST (for issuing)

// authorize route method is GET (for authorization)

// protected routes
// func (r *Router) Group(middleware ...MiddlewareFunc) *Router {
// 	group := &Router{
// 		base: r.base,
// 		tree: r.tree,
// 	}
// 	group.Use(middleware...)
// 	return group
// }

// token/refresh route method is POST (for token refreshment)

// authorize route method is POST (for approval)

// authorize route method is DELETE (for denial)

// tokens route method is GET (for listing user tokens)

// tokens/{tokenId} route method is DELETE (for revoking user tokens)

// clients route method is GET (for listing user clients)
// clients route method is POST (for creating user clients)
// clients/{clientId} route method is PUT (for updating user clients) - only secret and redirect uri
// clients/{clientId} route method is DELETE (for deleting user clients)
// scopes route method is GET (for listing available scopes)
// personal-access-tokens route method is GET (for listing user tokens)
// personal-access-tokens route method is POST (for creating user tokens)
// personal-access-tokens/{tokenId} route method is DELETE (for revoking user tokens)
