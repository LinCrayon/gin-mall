package routers

import (
	api "gin-mall/api/v1"
	"gin-mall/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret")) // 创建Cookie存储的会话存储对象
	r.Use(sessions.Sessions("mysession", store))              //mysession会话名称 \存储会话数据
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static")) //静态文件服务
	v1 := r.Group("api/v1")
	{
		//用户操作
		v1.POST("user/register", api.UserRegister()) //用户注册
		v1.POST("user/login", api.UserLogin())       //用户登录

		v1.GET("carousels", api.ListCarousels) // 轮播图

		//商品操作
		v1.GET("products_list", api.ListProduct)       //获取商品列表
		v1.GET("product/show/:id", api.ShowProduct)    //获取商品详细信息
		v1.GET("product/imgs/:id", api.ListProductImg) //获取商品图片地址
		v1.GET("categories", api.ListCategories)       // 商品分类

		authed := v1.Group("/") //需要登录的操作
		authed.Use(middleware.JWT())
		{
			//用户更新
			authed.PUT("user/update", api.UserUpdate)    //修改昵称
			authed.POST("user/avatar", api.UploadAvatar) //上传头像

			authed.POST("user/send_email", api.SendEmail)   //邮件发送
			authed.POST("user/valid_email", api.ValidEmail) //验证邮件

			// 显示金额
			authed.POST("money", api.ShowMoney)

			//商品操作
			authed.POST("product/create", api.CreateProduct) //创建商品
			authed.POST("product/search", api.SearchProduct) //搜索商品

			//收藏夹操作
			authed.POST("favorites", api.CreateFavorites)       //创建收藏夹
			authed.GET("favorites", api.ShowFavorites)          //显示收藏夹
			authed.DELETE("favorites/:id", api.DeleteFavorites) //删除收藏夹

			//地址模块
			authed.POST("address", api.CreateAddress)       //创建地址
			authed.GET("address/:id", api.GetAddress)       //获取详细的地址
			authed.GET("address/list", api.ListAddress)     //地址列表
			authed.PUT("address/:id", api.UpdateAddress)    //更新地址
			authed.DELETE("address/:id", api.DeleteAddress) //删除地址

			//购物车模块
			authed.POST("carts/create", api.CreateCart) //创建购物车
			authed.GET("carts/list", api.ListCart)      //购物车列表
			authed.PUT("carts/:id", api.UpdateCart)     //更新购物车
			authed.DELETE("carts/:id", api.DeleteCart)  //删除购物车

			// 订单操作
			authed.POST("orders_create", api.CreateOrder)
			authed.GET("orders_list", api.ListOrders)
			authed.GET("orders_show/:id", api.ShowOrder)
			authed.DELETE("orders_delete/:id", api.DeleteOrder)

			//支付模块
			authed.POST("paydown", api.OrderPay)
		}

	}

	return r
}
