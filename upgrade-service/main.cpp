#include <bits/stdc++.h>

using namespace std;

// Uncomment the following line if you need it
// typedef long long ll;

int main() {
	long long  k,sum=0,minsum=0;
	unsigned int n,m,temp,sig; 
	cin.sync_with_stdio(false);
	queue<long long> q1;
	long long q2;
	cin >> n >> m >> k;
	for (unsigned int i = 1; i <= m; i++) {
		long long ai;
		cin >> ai;
		q1.push(ai);
		sum+=ai;    
	}
	temp=1;
	sig=1;
	minsum=sum;
	for (unsigned int j=m+1;j<=n;j++)
	{
		long long bi;
		temp++;
		cin >>bi;
		q1.push(bi);
		sum+=bi-q1.front();
		if (sum<minsum)
		{
			minsum=sum;
			sig=temp;
		}		
		q1.pop();	
	}
	if (minsum < m*k)
	{
		cout << sig<<' '<<sig+m-1<<endl;
	}
    else{
    	cout<<0<<' '<<0<<endl;
	}
	return 0;
}
