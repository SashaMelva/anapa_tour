PGDMP                      |            anapa_db    16.2    16.1 1    O           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            P           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            Q           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            R           1262    32777    anapa_db    DATABASE     j   CREATE DATABASE anapa_db WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'C';
    DROP DATABASE anapa_db;
                postgres    false            �            1259    32779    accounts    TABLE     �   CREATE TABLE public.accounts (
    id integer NOT NULL,
    login character varying NOT NULL,
    password character varying NOT NULL,
    role character varying
);
    DROP TABLE public.accounts;
       public         heap    postgres    false            �            1259    32778    accounts_id_seq    SEQUENCE     �   ALTER TABLE public.accounts ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.accounts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    216            �            1259    32815    actions    TABLE       CREATE TABLE public.actions (
    id integer NOT NULL,
    name character varying NOT NULL,
    description text,
    category_id integer,
    date_start time with time zone,
    date_end time with time zone,
    activ boolean,
    organization_id integer
);
    DROP TABLE public.actions;
       public         heap    postgres    false            �            1259    32814    action_id_seq    SEQUENCE     �   ALTER TABLE public.actions ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.action_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    223            �            1259    32862    category_action    TABLE     f   CREATE TABLE public.category_action (
    id integer NOT NULL,
    name character varying NOT NULL
);
 #   DROP TABLE public.category_action;
       public         heap    postgres    false            �            1259    32861    category_action_id_seq    SEQUENCE     �   ALTER TABLE public.category_action ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.category_action_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    230            �            1259    32849    category_organization    TABLE     l   CREATE TABLE public.category_organization (
    id integer NOT NULL,
    name character varying NOT NULL
);
 )   DROP TABLE public.category_organization;
       public         heap    postgres    false            �            1259    32848    category_organization_id_seq    SEQUENCE     �   ALTER TABLE public.category_organization ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.category_organization_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    228            �            1259    32822    list_action_pins    TABLE     j   CREATE TABLE public.list_action_pins (
    id_action integer NOT NULL,
    id_pointer integer NOT NULL
);
 $   DROP TABLE public.list_action_pins;
       public         heap    postgres    false            �            1259    32826    organization    TABLE     �   CREATE TABLE public.organization (
    id integer NOT NULL,
    name character varying NOT NULL,
    description text,
    admin_id integer,
    legal_address character varying,
    full_name character varying,
    category_id integer
);
     DROP TABLE public.organization;
       public         heap    postgres    false            �            1259    32825    organization_id_seq    SEQUENCE     �   ALTER TABLE public.organization ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.organization_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    226            �            1259    32795    pins    TABLE     �   CREATE TABLE public.pins (
    id integer NOT NULL,
    name character varying NOT NULL,
    description text,
    activ boolean,
    organization_id integer,
    coordinat numeric[]
);
    DROP TABLE public.pins;
       public         heap    postgres    false            �            1259    32794    pins_id_seq    SEQUENCE     �   ALTER TABLE public.pins ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.pins_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    220            �            1259    32787    tokens    TABLE     u   CREATE TABLE public.tokens (
    id integer NOT NULL,
    accoun_id integer NOT NULL,
    token character varying
);
    DROP TABLE public.tokens;
       public         heap    postgres    false            �            1259    32786    tokens_id_seq    SEQUENCE     �   ALTER TABLE public.tokens ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.tokens_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    218            �            1259    32807    users    TABLE     �   CREATE TABLE public.users (
    id integer NOT NULL,
    first_name character varying NOT NULL,
    middel_name character varying NOT NULL,
    last_name character varying,
    date_birth time without time zone
);
    DROP TABLE public.users;
       public         heap    postgres    false            >          0    32779    accounts 
   TABLE DATA           =   COPY public.accounts (id, login, password, role) FROM stdin;
    public          postgres    false    216   8       E          0    32815    actions 
   TABLE DATA           s   COPY public.actions (id, name, description, category_id, date_start, date_end, activ, organization_id) FROM stdin;
    public          postgres    false    223   a8       L          0    32862    category_action 
   TABLE DATA           3   COPY public.category_action (id, name) FROM stdin;
    public          postgres    false    230   9       J          0    32849    category_organization 
   TABLE DATA           9   COPY public.category_organization (id, name) FROM stdin;
    public          postgres    false    228   89       F          0    32822    list_action_pins 
   TABLE DATA           A   COPY public.list_action_pins (id_action, id_pointer) FROM stdin;
    public          postgres    false    224   �9       H          0    32826    organization 
   TABLE DATA           n   COPY public.organization (id, name, description, admin_id, legal_address, full_name, category_id) FROM stdin;
    public          postgres    false    226   �9       B          0    32795    pins 
   TABLE DATA           X   COPY public.pins (id, name, description, activ, organization_id, coordinat) FROM stdin;
    public          postgres    false    220   ~:       @          0    32787    tokens 
   TABLE DATA           6   COPY public.tokens (id, accoun_id, token) FROM stdin;
    public          postgres    false    218   ;       C          0    32807    users 
   TABLE DATA           S   COPY public.users (id, first_name, middel_name, last_name, date_birth) FROM stdin;
    public          postgres    false    221   �;       S           0    0    accounts_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.accounts_id_seq', 9, true);
          public          postgres    false    215            T           0    0    action_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.action_id_seq', 5, true);
          public          postgres    false    222            U           0    0    category_action_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.category_action_id_seq', 3, true);
          public          postgres    false    229            V           0    0    category_organization_id_seq    SEQUENCE SET     J   SELECT pg_catalog.setval('public.category_organization_id_seq', 5, true);
          public          postgres    false    227            W           0    0    organization_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('public.organization_id_seq', 4, true);
          public          postgres    false    225            X           0    0    pins_id_seq    SEQUENCE SET     9   SELECT pg_catalog.setval('public.pins_id_seq', 4, true);
          public          postgres    false    219            Y           0    0    tokens_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.tokens_id_seq', 9, true);
          public          postgres    false    217            �           2606    32785    accounts accounts_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.accounts DROP CONSTRAINT accounts_pkey;
       public            postgres    false    216            �           2606    32821    actions action_pkey 
   CONSTRAINT     Q   ALTER TABLE ONLY public.actions
    ADD CONSTRAINT action_pkey PRIMARY KEY (id);
 =   ALTER TABLE ONLY public.actions DROP CONSTRAINT action_pkey;
       public            postgres    false    223            �           2606    32868 $   category_action category_action_pkey 
   CONSTRAINT     b   ALTER TABLE ONLY public.category_action
    ADD CONSTRAINT category_action_pkey PRIMARY KEY (id);
 N   ALTER TABLE ONLY public.category_action DROP CONSTRAINT category_action_pkey;
       public            postgres    false    230            �           2606    32855 0   category_organization category_organization_pkey 
   CONSTRAINT     n   ALTER TABLE ONLY public.category_organization
    ADD CONSTRAINT category_organization_pkey PRIMARY KEY (id);
 Z   ALTER TABLE ONLY public.category_organization DROP CONSTRAINT category_organization_pkey;
       public            postgres    false    228            �           2606    32832    organization organization_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.organization
    ADD CONSTRAINT organization_pkey PRIMARY KEY (id);
 H   ALTER TABLE ONLY public.organization DROP CONSTRAINT organization_pkey;
       public            postgres    false    226            �           2606    32801    pins pins_pkey 
   CONSTRAINT     L   ALTER TABLE ONLY public.pins
    ADD CONSTRAINT pins_pkey PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.pins DROP CONSTRAINT pins_pkey;
       public            postgres    false    220            �           2606    32793    tokens tokens_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT tokens_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.tokens DROP CONSTRAINT tokens_pkey;
       public            postgres    false    218            �           2606    32813    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    221            �           2606    32869    actions action_category_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.actions
    ADD CONSTRAINT action_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.category_action(id) NOT VALID;
 I   ALTER TABLE ONLY public.actions DROP CONSTRAINT action_category_id_fkey;
       public          postgres    false    223    3496    230            �           2606    32874 #   actions action_organization_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.actions
    ADD CONSTRAINT action_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organization(id) NOT VALID;
 M   ALTER TABLE ONLY public.actions DROP CONSTRAINT action_organization_id_fkey;
       public          postgres    false    3492    223    226            �           2606    32843 '   organization organization_admin_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.organization
    ADD CONSTRAINT organization_admin_id_fkey FOREIGN KEY (admin_id) REFERENCES public.accounts(id) NOT VALID;
 Q   ALTER TABLE ONLY public.organization DROP CONSTRAINT organization_admin_id_fkey;
       public          postgres    false    3482    226    216            �           2606    32856 *   organization organization_category_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.organization
    ADD CONSTRAINT organization_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.category_organization(id) NOT VALID;
 T   ALTER TABLE ONLY public.organization DROP CONSTRAINT organization_category_id_fkey;
       public          postgres    false    228    226    3494            �           2606    32838    tokens tokens_accoun_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT tokens_accoun_id_fkey FOREIGN KEY (accoun_id) REFERENCES public.accounts(id) NOT VALID;
 F   ALTER TABLE ONLY public.tokens DROP CONSTRAINT tokens_accoun_id_fkey;
       public          postgres    false    3482    218    216            >   2   x�3���O���,H,..�/J���2��!Sr3�,����1z\\\ �(�      E   �   x����	�0D��)4A�v�F'� I���� ��-N	�nV8mT9PH����9�F��3dmM� c�o³1��Fq�6z��<է2gr����Θ��vK�����p��ީ�YMF�^���Q{��*�������3���>0��w�      L   &   x�3�Lʯ�2�,�,�LN��2�L.-.������� �R      J   w   x�-�A
�@�;��Q����!��у��%�D]��B͏l�����&p��{?)~b��WŤ"�&p#��d����N�w���w���gk<�ɪ���h}�*>�;
���a��k3���UB      F      x������ � �      H   �   x�U���0E��� )�q�C.�+��V*i��7�N��ʱ�(�=����	ot�{d	蝜�I-g�:Ed��t2�M84���J�r�P޾��%̤!r]V�cnR�u����yZ9�:����},�a��=w��Մ��      B   }   x�u�A
�0����)r 	��h�����J�Q,B�^ndP(u������0c��������������E��|�U,�4y�����Җ:!G���~q5�_���Y8|�H�X�����P�      @   y   x�3�4�,/L-L-�2B0�L3ӒӒ3��+#�=9�?��3����/ӳ�3/�4����3� "���R/��@�D��%985�ѹĢ�$3��9%�"2#�83*1<�,�4�;8$��1���+F��� 	)*      C      x������ � �     