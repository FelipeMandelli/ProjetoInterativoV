from tkinter import *
from tkinter import ttk
from tkinter.ttk import Combobox
import requests
import json

import tkinter as tk

def salvar_nome_botao(nome):
    # Esta função será chamada quando um botão for pressionado
    # Ela salva o nome do botão na variável global nome_botao_escolhido
    global cadastro
    cadastro = nome
    print(f"Botão escolhido: {cadastro}")
    janela.destroy()

# Criar a janela principal
janela = tk.Tk()
fontePadrao = ('Open Sans Light', 10) #Fontes das labels
fontePadrao1 = ('Open Sans Light', 11, 'bold') #Fonte botao salvar
fontePadrao2 = ('Open Sans Light', 13, 'bold') #Fonte titulo
corFonte = '#494949' #Cor da fonte geral
janela.title('CHAMADA AUTOMÁTICA | PI V')

try:
    janela.iconbitmap(default='icones\\logo.ico')
except:
    pass

janela.resizable(False, False)
winHeight = 80
winWidth = 300
screen_width = janela.winfo_screenwidth()
screen_height = janela.winfo_screenheight()
x = (screen_width // 2) - (winWidth // 2)
y = (screen_height // 2) - (winHeight // 2)
janela.geometry(f'{winWidth}x{winHeight}+{x}+{y}')

frame_1 = Frame(janela, bd = 3, bg = 'white')
frame_1.place(relx = 0, rely=0, relwidth= 1, relheight= 0.24)
frame_2 = Frame(janela, bd = 3, bg = 'white')
frame_2.place(relx = 0, rely=0.24, relwidth= 1, relheight= 0.86)

lblTitle = Label(frame_1, text='\nSelecione o Cadastro\n', bg= 'white', font= fontePadrao2)
lblTitle.place(relx = 0, rely=0.02, relwidth= 1, relheight= 1)

# Variável para armazenar o nome do botão escolhido
cadastro = None

# Criar botões e associar a função salvar_nome_botao a cada botão
botao1 = tk.Button(frame_2, text="Aluno", bg= '#7404e3', bd=1, fg='white', font= fontePadrao1, command=lambda: salvar_nome_botao("Aluno"))
botao1.place(relx = 0.03, rely=0.05, relwidth= 0.29, relheight= 0.75)
botao2 = tk.Button(frame_2, text="Professor", bg= '#7404e3', bd=1, fg='white', font= fontePadrao1, command=lambda: salvar_nome_botao("Professor"))
botao2.place(relx = 0.35, rely=0.05, relwidth= 0.29, relheight= 0.75)
botao3 = tk.Button(frame_2, text="Matéria", bg= '#7404e3', bd=1, fg='white', font= fontePadrao1, command=lambda: salvar_nome_botao("Materia"))
botao3.place(relx = 0.67, rely=0.05, relwidth= 0.29, relheight= 0.75)

# Iniciar o loop principal da interface gráfica
janela.mainloop()

class FormPw:
    def __init__(self, title='Ficha Técnica | Supply Chain', txtWidth=32):
        if cadastro == "Aluno":
            win = Tk()
            fontePadrao = ('Open Sans Light', 10) #Fontes das labels
            fontePadrao1 = ('Open Sans Light', 11, 'bold') #Fonte botao salvar
            fontePadrao2 = ('Open Sans Light', 13, 'bold') #Fonte titulo
            corFonte = '#494949' #Cor da fonte geral
            win.title('CHAMADA AUTOMÁTICA | PI V')
            
            win.resizable(False, False)
            winHeight = 280
            winWidth = 480
            screen_width = win.winfo_screenwidth()
            screen_height = win.winfo_screenheight()
            x = (screen_width // 2) - (winWidth // 2)
            y = (screen_height // 2) - (winHeight // 2)
            win.geometry(f'{winWidth}x{winHeight}+{x}+{y}')
            
            img = PhotoImage(file='icones\\fotoPI.png')
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

            lblTitle = Label(frame_1, text='\nCadastro de Novo Aluno\n', bg= 'white', font= fontePadrao2)
            lblTitle.place(relx = 0, rely=0.02, relwidth= 1, relheight= 1)

            lbluserName = Label(frame_2, text=('Nome:'), bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lbluserName.place(relx = 0.02, rely=0.03, relwidth= 0.2, relheight= 0.1)
            txtuserName = Entry(frame_2, width=txtWidth)
            txtuserName.place(relx = 0.26, rely=0.03, relwidth=0.7, relheight= 0.1)

            lbluserEmail = Label(frame_2, text='E-mail:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lbluserEmail.place(relx = 0.02, rely=0.17, relwidth= 0.2, relheight= 0.1)
            txtuserEmail = Entry(frame_2, width=txtWidth)
            txtuserEmail.place(relx = 0.26, rely=0.17, relwidth=0.7, relheight= 0.1) 

            self.selected_option = StringVar(value="")
            style = ttk.Style()
            style.configure("White.TCheckbutton", background="white", foreground=corFonte, font=fontePadrao)

            lblPerfil = Label(frame_2, text='Perfil:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lblPerfil.place(relx = 0.02, rely=0.3, relwidth= 0.20, relheight= 0.1)
            checkOpcao1 = ttk.Checkbutton(frame_2, text="Professor", variable=self.selected_option, onvalue="P", style="White.TCheckbutton")
            checkOpcao1.place(relx=0.25, rely=0.29, relwidth=0.29, relheight=0.1)
            checkOpcao2 = ttk.Checkbutton(frame_2, text="Aluno", variable=self.selected_option, onvalue="S", style="White.TCheckbutton")
            checkOpcao2.place(relx=0.55, rely=0.29, relwidth=0.29, relheight=0.1)

            lblCartao = Label(frame_2, text='ID:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lblCartao.place(relx = 0.02, rely=0.43, relwidth= 0.20, relheight= 0.1)
            txtCartao = Entry(frame_2, width=txtWidth)
            txtCartao.place(relx = 0.26, rely=0.43, relwidth=0.7, relheight= 0.1)

            # lblCurso = Label(frame_2, text='Curso:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            # lblCurso.place(relx = 0.02, rely=0.48, relwidth= 0.2, relheight= 0.08)
            # txtCurso = Combobox(frame_2, values=['Engenharia da Computação', 'Engenharia de Produção'])
            # txtCurso.place(relx = 0.26, rely=0.48, relwidth=0.7, relheight= 0.08)
            # txtCurso.bind('<Key>', lambda e: 'break')

            lblCurso = Label(frame_2, text='Curso:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lblCurso.place(relx = 0.02, rely=0.71, relwidth= 0.24, relheight= 0.1)
            txtCurso = Combobox(frame_2, values=['Engenharia da Computação', 'Engenharia de Produção'])
            txtCurso.place(relx = 0.26, rely=0.71, relwidth=0.7, relheight= 0.1)
            txtCurso.bind('<Key>', lambda e: 'break')

            lblDocument = Label(frame_2, text='Telefone:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lblDocument.place(relx = 0.02, rely=0.57, relwidth= 0.2, relheight= 0.1)
            txtDocument = Entry(frame_2, width=txtWidth)
            txtDocument.place(relx = 0.26, rely=0.57, relwidth=0.7, relheight= 0.1)

            # lblTel = Label(frame_2, text='Telefone:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            # lblTel.place(relx = 0.02, rely=0.72, relwidth= 0.2, relheight= 0.08)
            # txtTel = Entry(frame_2, width=txtWidth)
            # txtTel.place(relx = 0.26, rely=0.72, relwidth=0.7, relheight= 0.08)

            btnOk = Button(frame_2, text = ('CADASTRAR'), bg= '#7404e3', bd=1, fg='white', font= fontePadrao1, command=lambda: self.btnOk_click(win, txtuserName, txtuserEmail, txtCurso, txtCartao, txtDocument))
            btnOk.place(relx = 0.02, rely=0.83, relwidth= 0.95, relheight= 0.16)

            win.mainloop()

        elif cadastro == "Professor":
            win = Tk()
            fontePadrao = ('Open Sans Light', 10) #Fontes das labels
            fontePadrao1 = ('Open Sans Light', 11, 'bold') #Fonte botao salvar
            fontePadrao2 = ('Open Sans Light', 13, 'bold') #Fonte titulo
            corFonte = '#494949' #Cor da fonte geral
            win.title('CHAMADA AUTOMÁTICA | PI V')
            
            win.resizable(False, False)
            winHeight = 250
            winWidth = 480
            screen_width = win.winfo_screenwidth()
            screen_height = win.winfo_screenheight()
            x = (screen_width // 2) - (winWidth // 2)
            y = (screen_height // 2) - (winHeight // 2)
            win.geometry(f'{winWidth}x{winHeight}+{x}+{y}')
            
            img = PhotoImage(file='icones\\fotoPI.png')
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

            lblTitle = Label(frame_1, text='\nCadastro de Novo Professor\n', bg= 'white', font= fontePadrao2)
            lblTitle.place(relx = 0, rely=0.02, relwidth= 1, relheight= 1)

            lbluserName = Label(frame_2, text=('Nome:'), bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lbluserName.place(relx = 0.02, rely=0.03, relwidth= 0.2, relheight= 0.11)
            txtuserName = Entry(frame_2, width=txtWidth)
            txtuserName.place(relx = 0.26, rely=0.03, relwidth=0.7, relheight= 0.11)

            lbluserEmail = Label(frame_2, text='E-mail:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lbluserEmail.place(relx = 0.02, rely=0.19, relwidth= 0.2, relheight= 0.11)
            txtuserEmail = Entry(frame_2, width=txtWidth)
            txtuserEmail.place(relx = 0.26, rely=0.19, relwidth=0.7, relheight= 0.11) 

            self.selected_option = StringVar(value="")
            style = ttk.Style()
            style.configure("White.TCheckbutton", background="white", foreground=corFonte, font=fontePadrao)

            lblPerfil = Label(frame_2, text='Perfil:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lblPerfil.place(relx = 0.02, rely=0.35, relwidth= 0.20, relheight= 0.11)
            checkOpcao1 = ttk.Checkbutton(frame_2, text="Professor", variable=self.selected_option, onvalue="P", style="White.TCheckbutton")
            checkOpcao1.place(relx=0.25, rely=0.34, relwidth=0.29, relheight=0.11)
            checkOpcao2 = ttk.Checkbutton(frame_2, text="Aluno", variable=self.selected_option, onvalue="S", style="White.TCheckbutton")
            checkOpcao2.place(relx=0.55, rely=0.34, relwidth=0.29, relheight=0.11)

            lblCartao = Label(frame_2, text='ID:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lblCartao.place(relx = 0.02, rely=0.51, relwidth= 0.20, relheight= 0.11)
            txtCartao = Entry(frame_2, width=txtWidth)
            txtCartao.place(relx = 0.26, rely=0.51, relwidth=0.7, relheight= 0.11)

            lblDocument = Label(frame_2, text='Telefone:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lblDocument.place(relx = 0.02, rely=0.67, relwidth= 0.2, relheight= 0.11)
            txtDocument = Entry(frame_2, width=txtWidth)
            txtDocument.place(relx = 0.26, rely=0.67, relwidth=0.7, relheight= 0.11)

            btnOk = Button(frame_2, text = ('CADASTRAR'), bg= '#7404e3', bd=1, fg='white', font= fontePadrao1, command=lambda: self.btnOk_click2(win, txtuserName, txtuserEmail, txtCartao, txtDocument))
            btnOk.place(relx = 0.02, rely=0.82, relwidth= 0.95, relheight= 0.16)

            win.mainloop()
        
        if cadastro == "Materia":
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
            
            img = PhotoImage(file='icones\\fotoPI2.png')
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

            lblTitle = Label(frame_1, text='\nCadastro de Nova Matéria\n', bg= 'white', font= fontePadrao2)
            lblTitle.place(relx = 0, rely=0.02, relwidth= 1, relheight= 1)

            lbluserName = Label(frame_2, text=('Matéria:'), bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lbluserName.place(relx = 0.02, rely=0.03, relwidth= 0.2, relheight= 0.07)
            txtuserName = Entry(frame_2, width=txtWidth)
            txtuserName.place(relx = 0.26, rely=0.03, relwidth=0.7, relheight= 0.07)

            lbluserEmail = Label(frame_2, text='Semestre:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lbluserEmail.place(relx = 0.02, rely=0.13, relwidth= 0.2, relheight= 0.07)
            txtuserEmail = Entry(frame_2, width=txtWidth)
            txtuserEmail.place(relx = 0.26, rely=0.13, relwidth=0.7, relheight= 0.07)

            lblCurso = Label(frame_2, text='Ano:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lblCurso.place(relx = 0.02, rely=0.23, relwidth= 0.20, relheight= 0.07)
            txtCurso = Entry(frame_2, width=txtWidth)
            txtCurso.place(relx = 0.26, rely=0.23, relwidth=0.7, relheight= 0.07)

            lblCartao = Label(frame_2, text='Professor:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lblCartao.place(relx = 0.02, rely=0.33, relwidth= 0.20, relheight= 0.07)
            txtCartao = Entry(frame_2, width=txtWidth)
            txtCartao.place(relx = 0.26, rely=0.33, relwidth=0.7, relheight= 0.07)

            lblDocument = Label(frame_2, text='ID Prof:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lblDocument.place(relx = 0.02, rely=0.43, relwidth= 0.2, relheight= 0.07)
            txtDocument = Entry(frame_2, width=txtWidth)
            txtDocument.place(relx = 0.26, rely=0.43, relwidth=0.7, relheight= 0.07)

            lblStudents = Label(frame_2, text='ID Alunos:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lblStudents.place(relx = 0.02, rely=0.53, relwidth= 0.2, relheight= 0.07)
            txtStudents = Entry(frame_2, width=txtWidth)
            txtStudents.place(relx = 0.26, rely=0.53, relwidth=0.7, relheight= 0.07)

            lblDia = Label(frame_2, text='Dia Sem:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lblDia.place(relx = 0.02, rely=0.63, relwidth= 0.2, relheight= 0.07)
            txtDia = Entry(frame_2, width=txtWidth)
            txtDia.place(relx = 0.26, rely=0.63, relwidth=0.7, relheight= 0.07)

            lblHorario = Label(frame_2, text='Horário:', bg= 'white', anchor=W, fg=corFonte, font=fontePadrao)
            lblHorario.place(relx = 0.02, rely=0.73, relwidth= 0.2, relheight= 0.07)
            txtHorario = Entry(frame_2, width=txtWidth)
            txtHorario.place(relx = 0.26, rely=0.73, relwidth=0.7, relheight= 0.07)

            btnOk = Button(frame_2, text = ('CADASTRAR'), bg= '#7404e3', bd=1, fg='white', font= fontePadrao1, command=lambda: self.btnOk_click3(win, txtuserName, txtuserEmail, txtCurso, txtCartao, txtDocument, txtStudents, txtDia, txtHorario))
            btnOk.place(relx = 0.02, rely=0.83, relwidth= 0.95, relheight= 0.16)

            win.mainloop()

    def btnOk_click(self, win, txtuserName, txtuserEmail, txtCurso, txtCartao, txtDocument):
        self.__userName = txtuserName.get()
        self.__userEmail = txtuserEmail.get()
        self.__curso = txtCurso.get()
        self.__cartao = txtCartao.get()
        self.__doc = txtDocument.get()

        api_url = "http://18.228.222.45:9015/new/registry"  # Substitua pelo URL da sua API
        data = {
            "name": self.__userName,
            "mail": self.__userEmail,
            "role": self.get_selected_option(),
            "course": self.__curso,
            "id": self.__cartao,
            "cellphone": self.__doc
        }
        response = requests.post(api_url, data=json.dumps(data))

        if response.status_code == 200:
            print("Chamada HTTP bem-sucedida")
        else:
            print("Erro na chamada HTTP:", response.status_code)

        win.destroy()

    def btnOk_click2(self, win, txtuserName, txtuserEmail, txtCartao, txtDocument):
        self.__userName = txtuserName.get()
        self.__userEmail = txtuserEmail.get()
        self.__cartao = txtCartao.get()
        self.__doc = txtDocument.get()

        api_url = "http://18.228.222.45:9015/new/registry"  # Substitua pelo URL da sua API
        data = {
            "name": self.__userName,
            "mail": self.__userEmail,
            "role": self.get_selected_option(),
            "id": self.__cartao,
            "cellphone": self.__doc
        }
        response = requests.post(api_url, data=json.dumps(data))

        if response.status_code == 200:
            print("Chamada HTTP bem-sucedida")
        else:
            print("Erro na chamada HTTP:", response.status_code)

        win.destroy()

    def btnOk_click3(self, win, txtuserName, txtuserEmail, txtCurso, txtCartao, txtDocument, txtStudents, txtDia, txtHorario):
        self.__userName = txtuserName.get()
        self.__userEmail = txtuserEmail.get()
        self.__curso = txtCurso.get()
        self.__cartao = txtCartao.get()
        self.__doc = txtDocument.get()
        self.__student = txtStudents.get()
        self.__dia = txtDia.get()
        self.__horario = txtHorario.get()


        api_url = "http://18.228.222.45:9015/new/subject"  # Substitua pelo URL da sua API
        data = {
            "subject_name": self.__userName,
            "reference_semester": int(self.__userEmail),
            "reference_year": self.__curso,
            "professor_name": self.__cartao,
            "professor_id": self.__doc,
            "students_enrolled_ids": self.__student,
            "weekday": int(self.__dia),
            "schedule": int(self.__horario)
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

    @property
    def student(self):
        try: return self.__student
        except: pass
    
    @property
    def dia(self):
        try: return self.__dia
        except: pass

    @property
    def horario(self):
        try: return self.__horario
        except: pass

if __name__ == "__main__":
    form_pw = FormPw()