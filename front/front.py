from tkinter import *
from tkinter import ttk
from tkinter.ttk import Combobox
import requests
import json

class FormPw:
    def __init__(self, title='Ficha Técnica | Supply Chain', txtWidth=32):
        win = Tk()
        fontePadrao = ('Open Sans Light', 10) #Fontes das labels
        fontePadrao1 = ('Open Sans Light', 11, 'bold') #Fonte botao salvar
        fontePadrao2 = ('Open Sans Light', 13, 'bold') #Fonte titulo
        corFonte = '#494949' #Cor da fonte geral
        win.title('CHAMADA AUTOMÁTICA | PI V')
        
        win.resizable(False, False)
        winHeight = 350
        winWidth = 480
        screen_width = win.winfo_screenwidth()
        screen_height = win.winfo_screenheight()
        x = (screen_width // 2) - (winWidth // 2)
        y = (screen_height // 2) - (winHeight // 2)
        win.geometry(f'{winWidth}x{winHeight}+{x}+{y}')
        
        img = PhotoImage(file='icones\\bg_sanofi.png')
        img=img.subsample(1,1)
        imgSanofi = Label(win, image=img)
        imgSanofi.place(relx= -0.005, rely= 0)
        try:
            win.iconbitmap(default='icones\\logo.ico')
        except:
            pass

        # Definição dos Frames na win
        frame_1 = Frame(win, bd = 3, bg = 'white')
        frame_1.place(relx = 0.3468, rely=0, relwidth= 0.6531, relheight= 0.14)
        frame_2 = Frame(win, bd = 3, bg = 'white')
        frame_2.place(relx = 0.3468, rely=0.14, relwidth= 0.6531, relheight= 0.86)

        lblTitle = Label(frame_1, text='\nCadastro de Novo Usuário\n', bg= 'white', font= fontePadrao2)
        lblTitle.place(relx = 0, rely=0.02, relwidth= 1, relheight= 1)

        lbluserName = Label(frame_2, text=('Nome:'), bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
        lbluserName.place(relx = 0.02, rely=0.03, relwidth= 0.2, relheight= 0.08)
        txtuserName = Entry(frame_2, width=txtWidth)
        txtuserName.place(relx = 0.26, rely=0.03, relwidth=0.7, relheight= 0.08)

        lbluserEmail = Label(frame_2, text='E-mail:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
        lbluserEmail.place(relx = 0.02, rely=0.14, relwidth= 0.2, relheight= 0.08)
        txtuserEmail = Entry(frame_2, width=txtWidth)
        txtuserEmail.place(relx = 0.26, rely=0.14, relwidth=0.7, relheight= 0.08) 

        self.selected_option = StringVar(value="")
        style = ttk.Style()
        style.configure("White.TCheckbutton", background="white", foreground=corFonte, font=fontePadrao)

        lblPerfil = Label(frame_2, text='Perfil:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
        lblPerfil.place(relx = 0.02, rely=0.26, relwidth= 0.20, relheight= 0.08)
        checkOpcao1 = ttk.Checkbutton(frame_2, text="Professor", variable=self.selected_option, onvalue="Professor", style="White.TCheckbutton")
        checkOpcao1.place(relx=0.25, rely=0.25, relwidth=0.29, relheight=0.08)
        checkOpcao2 = ttk.Checkbutton(frame_2, text="Aluno", variable=self.selected_option, onvalue="Aluno", style="White.TCheckbutton")
        checkOpcao2.place(relx=0.56, rely=0.25, relwidth=0.2, relheight=0.08)

        lblCartao = Label(frame_2, text='Cartão:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
        lblCartao.place(relx = 0.02, rely=0.36, relwidth= 0.20, relheight= 0.08)
        txtCartao = Entry(frame_2, width=txtWidth)
        txtCartao.place(relx = 0.26, rely=0.36, relwidth=0.7, relheight= 0.08)

        lblCurso = Label(frame_2, text='Curso:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
        lblCurso.place(relx = 0.02, rely=0.48, relwidth= 0.2, relheight= 0.08)
        txtCurso = Combobox(frame_2, values=['Engenharia da Computação', 'Engenharia de Produção'])
        txtCurso.place(relx = 0.26, rely=0.48, relwidth=0.7, relheight= 0.08)
        txtCurso.bind('<Key>', lambda e: 'break')

        lblDocument = Label(frame_2, text='Documento:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
        lblDocument.place(relx = 0.02, rely=0.6, relwidth= 0.24, relheight= 0.08)
        txtDocument = Entry(frame_2, width=txtWidth)
        txtDocument.place(relx = 0.26, rely=0.6, relwidth=0.7, relheight= 0.08)

        lblTel = Label(frame_2, text='Telefone:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
        lblTel.place(relx = 0.02, rely=0.72, relwidth= 0.2, relheight= 0.08)
        txtTel = Entry(frame_2, width=txtWidth)
        txtTel.place(relx = 0.26, rely=0.72, relwidth=0.7, relheight= 0.08)

        btnOk = Button(frame_2, text = ('CADASTRAR'), bg= '#7404e3', bd=1, fg='white', font= fontePadrao1, command=lambda: self.btnOk_click(win, txtuserName, txtuserEmail, txtCurso, txtCartao, txtDocument, txtTel))
        btnOk.place(relx = 0.02, rely=0.84, relwidth= 0.95, relheight= 0.14)

        win.mainloop()

    def btnOk_click(self, win, txtuserName, txtuserEmail, txtCurso, txtCartao, txtDocument, txtTel):
        self.__userName = txtuserName.get()
        self.__userEmail = txtuserEmail.get()
        self.__curso = txtCurso.get()
        self.__cartao = txtCartao.get()
        self.__doc = txtDocument.get()
        self.__tel = txtTel.get()

        # # Fazer a chamada HTTP para a API
        api_url = "http://localhost:9015/new/registry"  # Substitua pelo URL da sua API
        data = {
            "Name": self.__userName,
            "Mail": self.__userEmail,
            "Role": self.get_selected_option(),
            "Course": self.__curso,
            "Tag": self.__cartao,
            "Document": self.__doc,
            "Tel": self.__tel
        }
        response = requests.post(api_url, data=json.dumps(data))

        if response.status_code == 200:
            print("Chamada HTTP bem-sucedida")
        else:
            print("Erro na chamada HTTP:", response.status_code)

        win.destroy()

    def get_selected_option(self):
        selected_option = self.selected_option.get()
        return selected_option


    def btnCancel_click(self, win):
        win.destroy()

    @property
    def userName(self):
        try: return self.__userName
        except: pass
    
    @property
    def userEmail(self):
        try: return self.__userEmail
        except: pass

    @property
    def curso(self):
        try: return self.__curso
        except: pass

    @property
    def cartao(self):
        try: return self.__cartao
        except: pass
    
    @property
    def doc(self):
        try: return self.__doc
        except: pass

    @property
    def tel(self):
        try: return self.__tel
        except: pass

if __name__ == "__main__":
    form = FormPw()
    user_Name = form.userName
    user_Email = form.userEmail
    Curso = form.curso
    Cartao = form.cartao
    Perfil = form.get_selected_option()
    Documento = form.doc
    Telefone = form.tel
    print("Nome: {}" .format(user_Name))
    print("E-mail: {}" .format(user_Email))
    print("Perfil: {}" .format(Perfil))
    print("Cartão: {}".format(Cartao))
    print("Curso: {}".format(Curso))
    print("Documento: {}".format(Documento))
    print("Tel: {}".format(Telefone))