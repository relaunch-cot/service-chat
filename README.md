# Chat Service - Relaunch 

Serviço responsável por gerenciar todas as funcionalidades relacionadas ao sistema de chat da plataforma Relaunch.

## 📋 Sobre o Projeto

Este microserviço é parte da arquitetura da plataforma Relaunch e fornece funcionalidades completas de mensagens e chat entre usuários, incluindo:

- Criação e gerenciamento de conversas entre usuários
- Envio e recebimento de mensagens em tempo real
- Histórico de mensagens por conversa
- Listagem de chats ativos por usuário
- Busca e validação de conversas existentes

## 🛠️ Tecnologias Utilizadas

- **Go** - Linguagem de programação principal
- **MySQL** - Banco de dados relacional para persistência
- **gRPC** - Protocolo de comunicação entre serviços
- **Context** - Gerenciamento de requisições e timeouts

## 🏗️ Arquitetura

O serviço segue os princípios de Clean Architecture, com separação clara de responsabilidades:

- **Repositories**: Camada de acesso aos dados (MySQL)
- **Models**: Estruturas de dados compartilhadas
- **gRPC Status Codes**: Tratamento padronizado de erros

## 🚀 Funcionalidades Principais

- **Gerenciamento de Chats**: Criação de conversas entre dois usuários com validação de duplicidade
- **Sistema de Mensagens**: Envio de mensagens com validação de participação no chat
- **Consultas Otimizadas**: Recuperação eficiente de mensagens e conversas com JOINs
- **Tratamento de Erros**: Respostas padronizadas com códigos gRPC apropriados

## 💡 Destaques Técnicos

- Utilização de prepared statements para segurança
- Validação de permissões antes de operações críticas
- Queries otimizadas com índices apropriados
- Tratamento robusto de erros e edge cases