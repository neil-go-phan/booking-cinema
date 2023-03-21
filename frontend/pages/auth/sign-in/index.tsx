import { NextPage } from 'next'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faUser } from '@fortawesome/free-regular-svg-icons'
import { faLock } from '@fortawesome/free-solid-svg-icons'
import {
  Button, Col, Container, Form, InputGroup, Row,
} from 'react-bootstrap'
import Link from 'next/link'
import { useContext, useState } from 'react'
import { useRouter } from 'next/router'
import axios from 'axios'
import { setCookie } from 'cookies-next';
import { useForm, SubmitHandler } from 'react-hook-form';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import cryptoJS from 'crypto-js';
import { _REGEX } from '@/constants/regex'
import { _ROUTES } from '@/constants/route'
import { IsLoggedContext } from '@/helpers/isLoggedContext'
interface SignInFormProperty {
  username: string;
  password: string;
}

const SignIn: NextPage = () => {
  const router = useRouter()
  const [errorMessage, setErrorMessage] = useState({
    trigger: false,
    message: '',
  });
  const logged = useContext(IsLoggedContext)
  const schema = yup.object().shape({
    username: yup
      .string()
      .required('Username must not be empty')
      .min(8, 'Username must have 8-16 character')
      .max(16, 'Username must have 8-16 character')
      .matches(
        _REGEX.REGEX_USENAME_PASSWORD,
        'Username must not contain special character like @#$^...'
      ),
    password: yup
      .string()
      .required('Password must not be empty')
      .min(8, 'Password must have 8-16 character')
      .max(16, 'Password must have 8-16 character')
      .matches(
        _REGEX.REGEX_USENAME_PASSWORD,
        'Password must not contain special character like @#$^...'
      ),
  });
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<SignInFormProperty>({
    resolver: yupResolver(schema),
  });

  const onSubmit: SubmitHandler<SignInFormProperty> = async (data) => {
    let { username, password } = data;
    password = cryptoJS.SHA512(password).toString()
    try {
      const res = await axios.post(`${process.env.NEXT_PUBLIC_DOMAIN}auth/sign-in`, {
        username,
        password,
      });
      setCookie('access_token', res.data.accessToken);
      setCookie('refresh_token', res.data.refreshToken);
      setErrorMessage({
        trigger: false,
        message: res.data.message,
      });
      router.push(_ROUTES.HOME_PAGE);
      logged?.setIsLogged(true)
    } catch (error: any) {
      setErrorMessage({
        trigger: true,
        message: error.response.data.message,
      });
    }
    reset({ password: '' });

  };

  return (
    <div className='bg-light min-vh-100 d-flex flex-row align-items-center dark:bg-transparent auth'>
      <Container>
        <Row className='justify-content-center align-items-center px-3'>
          <Col lg={8}>
            <Row>
              <Col md={7} className='bg-white border p-5'>
                <div className=''>
                  <h1>Sign in</h1>
                  <p className='text-black-50'>Sign in to your account</p>

                  <form onSubmit={handleSubmit(onSubmit)}>
                    <InputGroup className='mb-3'>
                      <InputGroup.Text>
                        <FontAwesomeIcon
                          icon={faUser}
                          fixedWidth
                        />
                      </InputGroup.Text>
                      <Form.Control
                        {...register('username')}
                        placeholder='Type your username'
                        type='text'
                        required
                      />
                    </InputGroup>

                    {errors.username && (
                      <p className='errorMessage'>{errors.username.message}</p>
                    )}

                    <InputGroup className='mb-3'>
                      <InputGroup.Text>
                        <FontAwesomeIcon
                          icon={faLock}
                          fixedWidth
                        />
                      </InputGroup.Text>
                      <Form.Control
                        {...register('password')}
                        placeholder='Type your password'
                        type='password'
                        required
                      />
                    </InputGroup>

                    {errors.password && (
                      <p className='errorMessage'>{errors.password.message}</p>
                    )}

                    {errorMessage.trigger && (
                      <p className='errorMessage'>{errorMessage.message}</p>
                    )}

                    <Row>
                      <Col xs={6}>
                        <Button className='px-4' variant='primary' type='submit'>Login</Button>
                      </Col>
                      <Col xs={6} className='text-end'>
                        {/* <Button className='px-0' variant='link'>
                          Forgot
                          password?
                        </Button> */}
                      </Col>
                    </Row>
                  </form>
                </div>
              </Col>
              <Col
                md={5}
                className='bg-primary text-white d-flex align-items-center justify-content-center p-5'
              >
                <div className='text-center'>
                  <h2>Sign up</h2>
                  <p>
                    Signing up to experience all of our great features !
                  </p>
                  <Link href={_ROUTES.SIGN_UP_PAGE}>
                    <button className='btn btn-lg btn-outline-light mt-3' type='button'>
                      Register Now!
                    </button>
                  </Link>
                </div>
              </Col>
            </Row>
          </Col>
        </Row>
      </Container>
    </div>
  )
}

export default SignIn
