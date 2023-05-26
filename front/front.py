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
        winHeight = 300
        winWidth = 450
        screen_width = win.winfo_screenwidth()
        screen_height = win.winfo_screenheight()
        x = (screen_width // 2) - (winWidth // 2)
        y = (screen_height // 2) - (winHeight // 2)
        win.geometry(f'{winWidth}x{winHeight}+{x}+{y}')
        
        # img = PhotoImage(file='icones\\bg_sanofi.png')
        # img=img.subsample(1,1)
        # imgSanofi = Label(win, image=img)
        # imgSanofi.place(relx= -0.005, rely= 0)
        # try:
        #     win.iconbitmap(default='icones\\logo.ico')
        # except:
        #     pass

        # Definição dos Frames na win
        frame_1 = Frame(win, bd = 3, bg = 'white')
        frame_1.place(relx = 0.37, rely=0, relwidth= 0.63, relheight= 0.14)
        frame_2 = Frame(win, bd = 3, bg = 'white')
        frame_2.place(relx = 0.37, rely=0.14, relwidth= 0.63, relheight= 0.86)

        lblTitle = Label(frame_1, text='\nCadastro de Novo Usuário\n', bg= 'white', font= fontePadrao2)
        lblTitle.place(relx = 0, rely=0.05, relwidth= 1, relheight= 1)

        lbluserName = Label(frame_2, text=('Nome:'), bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
        lbluserName.place(relx = 0.05, rely=0.03, relwidth= 0.2, relheight= 0.10)
        txtuserName = Entry(frame_2, width=txtWidth)
        txtuserName.place(relx = 0.23, rely=0.03, relwidth= 0.72, relheight= 0.10)

        lbluserEmail = Label(frame_2, text='E-mail:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
        lbluserEmail.place(relx = 0.05, rely=0.19, relwidth= 0.2, relheight= 0.10)
        txtuserEmail = Entry(frame_2, width=txtWidth)
        txtuserEmail.place(relx = 0.23, rely=0.19, relwidth= 0.72, relheight= 0.10) 

        self.selected_option = StringVar(value="")
        style = ttk.Style()
        style.configure("White.TCheckbutton", background="white", foreground=corFonte, font=fontePadrao)

        lblPerfil = Label(frame_2, text='Perfil:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
        lblPerfil.place(relx = 0.05, rely=0.34, relwidth= 0.20, relheight= 0.10)
        checkOpcao1 = ttk.Checkbutton(frame_2, text="Professor", variable=self.selected_option, onvalue="Professor", style="White.TCheckbutton")
        checkOpcao1.place(relx=0.22, rely=0.34, relwidth=0.29, relheight=0.10)
        checkOpcao2 = ttk.Checkbutton(frame_2, text="Aluno", variable=self.selected_option, onvalue="Aluno", style="White.TCheckbutton")
        checkOpcao2.place(relx=0.53, rely=0.34, relwidth=0.2, relheight=0.10)

        lblCartao = Label(frame_2, text='Cartão:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
        lblCartao.place(relx = 0.05, rely=0.49, relwidth= 0.20, relheight= 0.10)
        txtCartao = Entry(frame_2, width=txtWidth)
        txtCartao.place(relx = 0.23, rely=0.49, relwidth= 0.72, relheight= 0.10)

        lblCurso = Label(frame_2, text='Curso:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
        lblCurso.place(relx = 0.05, rely=0.66, relwidth= 0.2, relheight= 0.10)
        txtCurso = Combobox(frame_2, values=['Engenharia da Computação', 'Engenharia de Produção'])
        txtCurso.place(relx = 0.23, rely=0.66, relwidth= 0.72, relheight= 0.10)
        txtCurso.bind('<Key>', lambda e: 'break') 

        btnOk = Button(frame_2, text = ('CADASTRAR'), bg= '#7404e3', bd=1, fg='white', font= fontePadrao1, command=lambda: self.btnOk_click(win, txtuserName, txtuserEmail, txtCurso, txtCartao))
        btnOk.place(relx = 0.05, rely=0.8, relwidth= 0.9, relheight= 0.17)

        win.mainloop()

    def btnOk_click(self, win, txtuserName, txtuserEmail, txtCurso, txtCartao):
        self.__userName = txtuserName.get()
        self.__userEmail = txtuserEmail.get()
        self.__curso = txtCurso.get()
        self.__cartao = txtCartao.get()

        # Fazer a chamada HTTP para a API
        api_url = "http://localhost:9015/new/registry"  # Substitua pelo URL da sua API
        data = {
            "Name": self.__userName,
            "Mail": self.__userEmail,
            "Role": self.get_selected_option(),
            "Course": self.__curso,
            "Tag": self.__cartao
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

if __name__ == "__main__":
    form = FormPw()
    user_Name = form.userName
    user_Email = form.userEmail
    Curso = form.curso
    Cartao = form.cartao
    Perfil = form.get_selected_option()
    print("Nome: {}" .format(user_Name))
    print("E-mail: {}" .format(user_Email))
    print("Perfil: {}" .format(Perfil))
    print("Cartão: {}".format(Cartao))
    print("Curso: {}".format(Curso))
