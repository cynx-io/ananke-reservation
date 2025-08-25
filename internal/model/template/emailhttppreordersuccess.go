package template

import "fmt"

const emailHttpPreorderSuccess = `
<!DOCTYPE html>
<html lang="id">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Pembayaran Preorder Berhasil</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background: linear-gradient(135deg, #000000 0%%, #1a1a1a 100%%);
            color: #ffffff;
            line-height: 1.6;
        }
        
        .container {
            max-width: 600px;
            margin: 0 auto;
            background: #000000;
            border: 2px solid #FFD700;
            border-radius: 20px;
            overflow: hidden;
            box-shadow: 0 20px 40px rgba(255, 215, 0, 0.2);
        }
        
        .header {
            background: linear-gradient(135deg, #FFD700 0%%, #FFA500 100%%);
            padding: 40px 30px;
            text-align: center;
            position: relative;
        }
        
        .header::before {
            content: '';
            position: absolute;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
            background: url('data:image/svg+xml,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100"><defs><pattern id="grid" width="10" height="10" patternUnits="userSpaceOnUse"><path d="M 10 0 L 0 0 0 10" fill="none" stroke="rgba(0,0,0,0.1)" stroke-width="0.5"/></pattern></defs><rect width="100" height="100" fill="url(%%23grid)" /></svg>');
            opacity: 0.3;
        }
        
        .header h1 {
            color: #000000;
            font-size: 32px;
            font-weight: 700;
            margin-bottom: 10px;
            text-shadow: 0 2px 4px rgba(0,0,0,0.1);
            position: relative;
            z-index: 1;
        }
        
        .success-icon {
            width: 80px;
            height: 80px;
            background: #000000;
            border-radius: 50%%;
            margin: 0 auto 20px;
            display: flex;
            align-items: center;
            justify-content: center;
            position: relative;
            z-index: 1;
        }
        
        .checkmark {
            width: 40px;
            height: 40px;
            border: 3px solid #FFD700;
            border-radius: 50%%;
            position: relative;
        }
        
        .checkmark::after {
            content: '';
            position: absolute;
            left: 12px;
            top: 6px;
            width: 8px;
            height: 16px;
            border: solid #FFD700;
            border-width: 0 3px 3px 0;
            transform: rotate(45deg);
        }
        
        .content {
            padding: 40px 30px;
            background: linear-gradient(180deg, #000000 0%%, #0a0a0a 100%%);
        }
        
        .success-message {
            text-align: center;
            margin-bottom: 30px;
        }
        
        .success-message h2 {
            color: #FFD700;
            font-size: 28px;
            margin-bottom: 15px;
            font-weight: 600;
        }
        
        .success-message p {
            color: #cccccc;
            font-size: 16px;
            margin-bottom: 10px;
        }
        
        .highlight-box {
            background: linear-gradient(135deg, #1a1a1a 0%%, #2a2a2a 100%%);
            border: 1px solid #FFD700;
            border-radius: 15px;
            padding: 25px;
            margin: 30px 0;
            text-align: center;
        }
        
        .highlight-box h3 {
            color: #FFD700;
            font-size: 20px;
            margin-bottom: 15px;
        }
        
        .highlight-box p {
            color: #ffffff;
            font-size: 14px;
            margin-bottom: 20px;
        }
        
        .discord-button {
            display: inline-block;
            background: linear-gradient(135deg, #7289DA 0%%, #5865F2 100%%);
            color: #ffffff;
            text-decoration: none;
            padding: 15px 30px;
            border-radius: 30px;
            font-weight: 600;
            font-size: 16px;
            transition: all 0.3s ease;
            box-shadow: 0 4px 15px rgba(114, 137, 218, 0.3);
            position: relative;
            overflow: hidden;
        }
        
        .discord-button::before {
            content: '';
            position: absolute;
            top: 0;
            left: -100%%;
            width: 100%%;
            height: 100%%;
            background: linear-gradient(90deg, transparent, rgba(255,255,255,0.2), transparent);
            transition: left 0.5s;
        }
        
        .discord-button:hover::before {
            left: 100%%;
        }
        
        .discord-button:hover {
            transform: translateY(-2px);
            box-shadow: 0 8px 25px rgba(114, 137, 218, 0.4);
        }
        
        .info-grid {
            display: grid;
            grid-template-columns: 1fr 1fr;
            gap: 20px;
            margin: 30px 0;
        }
        
        .info-item {
            background: #1a1a1a;
            border: 1px solid #333333;
            border-radius: 10px;
            padding: 20px;
            text-align: center;
        }
        
        .info-item h4 {
            color: #FFD700;
            font-size: 14px;
            margin-bottom: 8px;
            text-transform: uppercase;
            letter-spacing: 1px;
        }
        
        .info-item p {
            color: #ffffff;
            font-size: 16px;
            font-weight: 600;
        }
        
        .footer {
            background: #0a0a0a;
            padding: 30px;
            text-align: center;
            border-top: 1px solid #333333;
        }
        
        .footer p {
            color: #888888;
            font-size: 12px;
            margin-bottom: 10px;
        }
        
        .social-links {
            margin-top: 20px;
        }
        
        .social-links a {
            color: #FFD700;
            text-decoration: none;
            margin: 0 15px;
            font-size: 14px;
            transition: color 0.3s ease;
        }
        
        .social-links a:hover {
            color: #ffffff;
        }
        
        @media (max-width: 480px) {
            .container {
                margin: 10px;
                border-radius: 15px;
            }
            
            .header {
                padding: 30px 20px;
            }
            
            .header h1 {
                font-size: 24px;
            }
            
            .content {
                padding: 30px 20px;
            }
            
            .info-grid {
                grid-template-columns: 1fr;
            }
            
            .discord-button {
                padding: 12px 25px;
                font-size: 14px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <div class="success-icon">
                <div class="checkmark"></div>
            </div>
            <h1>Pembayaran Berhasil!</h1>
        </div>
        
        <div class="content">
            <div class="success-message">
                <h2>Preorder Anda Telah Dikonfirmasi</h2>
                <p>Terima kasih telah melakukan preorder dengan kami.</p>
                <p>Pembayaran Anda telah berhasil diproses dan pesanan sedang dipersiapkan.</p>
            </div>
            
            <div class="highlight-box">
                <h3>🎉 Selamat Datang di Komunitas Eksklusif!</h3>
                <p>Bergabunglah dengan Discord server kami untuk mendapatkan update terbaru, akses eksklusif, dan berinteraksi dengan komunitas premium lainnya.</p>
                <a href="%s" class="discord-button">
                    🎮 Bergabung ke Discord
                </a>
            </div>
            
            <div class="info-grid">
                <div class="info-item">
                    <h4>Status Pesanan</h4>
                    <p>Dikonfirmasi</p>
                </div>
                <div class="info-item">
                    <h4>Estimasi Pengiriman</h4>
                    <p>1-3 Hari Kerja</p>
                </div>
            </div>
            
            <div class="highlight-box">
                <h3>📧 Informasi Penting</h3>
                <p>Kami akan mengirimkan email update mengenai status pengiriman pesanan Anda. Pastikan untuk memeriksa folder spam jika tidak menerima email dari kami.</p>
            </div>
        </div>
        
        <div class="footer">
            <p>Email ini dikirim secara otomatis, mohon tidak membalas email ini.</p>
            <p>Jika ada pertanyaan, silakan hubungi tim support kami.</p>
            
            <div class="social-links">
                <a href="%s">Instagram</a>
                <a href="%s">TikTok</a>
                <a href="%s">Facebook</a>
                <a href="%s">Twitter</a>
                <a href="%s">Website</a>
            </div>
            
            <p style="margin-top: 20px; color: #FFD700;">
                ✨ Terima kasih telah mempercayai kami ✨
            </p>
        </div>
    </div>
</body>
</html>
`

func GenerateEmailHttpPreorderSuccess(links Links) string {
	return fmt.Sprintf(emailHttpPreorderSuccess, links.DiscordLink, links.InstagramLink, links.TiktokLink, links.FacebookLink, links.TwitterLink, links.WebsiteLink)
}
