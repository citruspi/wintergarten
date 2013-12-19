import os
import requests

class film (object):

    def __init__(self, tmdb_id, poster_path_base):

        response = requests.get('http://api.themoviedb.org/3/movie/'+str(tmdb_id), 
                            params={'api_key':os.environ['TMDB_API_KEY']}).json()

        self.adult = response['adult']
        self.backdrop_path = response['backdrop_path']
        self.budget = response['budget']
        self.genres = response['genres']
        self.homepage = response['homepage']
        self.tmdb_id = response['id']
        self.original_title = response['original_title']
        self.overview = response['overview']
        self.poster_path = response['poster_path']
        self.popularity = response['popularity']
        self.production_companies = response['production_companies']
        self.production_countries = response['production_countries']
        self.release_date = response['release_date']
        self.release_year = response['release_date'].split('-')[0]
        self.revenue = response['revenue']
        self.runtime = response['runtime']
        self.spoken_languages = response['spoken_languages']
        self.status = response['status']
        self.tagline = response['tagline']
        self.title = response['title']
        self.vote_average = response['vote_average']
        self.vote_count = response['vote_count']

        if self.poster_path is None:

        	self.poster_path = 'http://dummyimage.com/1000x1500/000000/ffffff.png&text=No+Poster+Found+:/'

        else:

        	self.poster_path = poster_path_base + 'original' + self.poster_path